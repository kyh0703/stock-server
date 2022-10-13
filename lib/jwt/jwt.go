package jwt

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kyh0703/stock-server/config"
)

type TokenInfo struct {
	AccessToken         string
	AccessUUID          string
	AccessTokenExpires  int64
	RefreshUUID         string
	RefreshToken        string
	RefreshTokenExpires int64
}

func generateToken(id int, expire int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id
	claims["exp"] = expire
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env.APISecret))
}

func CreateToken(userID int) (*TokenInfo, error) {
	t := new(TokenInfo)
	t.AccessTokenExpires = time.Now().Add(time.Minute * 15).Unix()
	t.AccessUUID = uuid.NewString()
	accessToken, err := generateToken(userID, t.AccessTokenExpires)
	if err != nil {
		return nil, err
	}
	t.AccessToken = accessToken
	t.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	t.RefreshUUID = uuid.NewString()
	refreshToken, err := generateToken(userID, t.RefreshTokenExpires)
	if err != nil {
		return nil, err
	}
	t.RefreshToken = refreshToken
	return t, nil
}

func ValidateTokenFromCookie(accessToken string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, &claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Env.APISecret), nil
		})
	return claims, err
}

func ValidateToken(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint32, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}
