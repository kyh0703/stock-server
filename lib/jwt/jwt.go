package jwt

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/kyh0703/stock-server/config"
)

type TokenMetaData struct {
	AccessToken         string
	AccessUUID          string
	AccessTokenExpires  int64
	RefreshUUID         string
	RefreshToken        string
	RefreshTokenExpires int64
}

type AccessData struct {
	AccessUUID string
	UserID     int
}

func generateAccessToken(id int, UUID string, expire int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["access_uuid"] = UUID
	claims["user_id"] = id
	claims["exp"] = expire
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env.APISecret))
}

func generateRefreshToken(id int, UUID string, expire int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["refresh_uuid"] = UUID
	claims["user_id"] = id
	claims["exp"] = expire
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env.APISecret))
}

func CreateToken(userID int) (*TokenMetaData, error) {
	t := new(TokenMetaData)
	t.AccessTokenExpires = time.Now().Add(time.Minute * 15).Unix()
	t.AccessUUID = uuid.NewString()
	accessToken, err := generateAccessToken(userID, t.AccessUUID, t.AccessTokenExpires)
	if err != nil {
		return nil, err
	}
	t.AccessToken = accessToken
	t.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	t.RefreshUUID = uuid.NewString()
	refreshToken, err := generateRefreshToken(userID, t.RefreshUUID, t.RefreshTokenExpires)
	if err != nil {
		return nil, err
	}
	t.RefreshToken = refreshToken
	return t, nil
}

func SaveTokenData(client *redis.Client, userID int, token *TokenMetaData) error {
	// convert Unis to UTC
	var (
		at  = time.Unix(token.AccessTokenExpires, 0)
		rt  = time.Unix(token.RefreshTokenExpires, 0)
		now = time.Now()
	)
	if err := client.Set(token.AccessUUID, strconv.Itoa(userID), at.Sub(now)).Err(); err != nil {
		return err
	}
	if err := client.Set(token.RefreshUUID, strconv.Itoa(userID), rt.Sub(now)).Err(); err != nil {
		return err
	}
	return nil
}

func DeleteTokenData(client *redis.Client, UUID string) (int64, error) {
	deleted, err := client.Del(UUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func ValidateTokenFromCookie(accessToken string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, &claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Env.APISecret), nil
		})
	return claims, err
}

func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ValidateToken(c *gin.Context) error {
	token, err := VerifyToken(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
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

func ExtractTokenMetadata(c *gin.Context) (*AccessData, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return nil, err
		}
		return &AccessData{
			AccessUUID: accessUUID,
			UserID:     int(userID),
		}, nil
	}
	return nil, errors.New("failed extract token metadata")
}
