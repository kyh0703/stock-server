package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kyh0703/stock-server/config"
)

func CreateToken(userId uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]bytes(config.Env.API_SECRET))
}
