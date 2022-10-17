package auth

import (
	"fmt"

	"github.com/kyh0703/stock-server/lib/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (svc *AuthService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash: %w", err)
	}
	return string(hash), nil
}

func (svc *AuthService) CompareHashPassword(hash, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		if err != bcrypt.ErrMismatchedHashAndPassword {
			return false, fmt.Errorf("mismatch password: %w", err)
		}
		return false, err
	}
	return true, nil
}

func (svc *AuthService) Login(id int) (map[string]string, error) {
	// create token
	token, err := jwt.CreateToken(id)
	if err != nil {
		return nil, err
	}
	// save the redis
	if err := jwt.SaveTokenData(id, token); err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}, nil
}

func (svc *AuthService) Verify(jwtString string) {
}
