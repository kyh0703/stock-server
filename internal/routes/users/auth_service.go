package users

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kyh0703/stock-server/configs"
	"golang.org/x/crypto/bcrypt"
)

const (
	TokenKeyAuthorized = "authorized"
	TokenKeyUserID     = "userID"
	TokenKeyExpire     = "expire"
)

type AuthService struct {
	authRepo AuthRepository
}

func (svc *AuthService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash: %w", err)
	}

	return string(hash), nil
}

func (svc *AuthService) CompareHashPassword(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err != bcrypt.ErrMismatchedHashAndPassword {
			return false, fmt.Errorf("mismatch password: %w", err)
		}
		return false, err
	}

	return true, nil
}

func (svc *AuthService) Login(userID int) (*dto.AccessTokenDto, error) {
	accessToken, refreshToken, err := svc.generateToken(userID)
	if err != nil {
		return nil, err
	}

	claims, err := svc.RefreshTokenData(refreshToken)
	if err != nil {
		return nil, err
	}

	expire, err := strconv.ParseInt(
		fmt.Sprintf("%.0f", claims[TokenKeyExpire]),
		10, 64)
	if err != nil {
		return nil, err
	}

	expireTime := time.Unix(expire, 0)
	err = svc.authRepo.InsertToken(userID, expireTime)
	if err != nil {
		return nil, err
	}

	return &dto.AccessTokenDto{
		Token:  accessToken,
		Expire: expireTime,
	}, nil
}

func (svc *AuthService) Refresh(refreshToken string) (*dto.AccessTokenDto, error) {
	claims, err := svc.RefreshTokenData(refreshToken)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.ParseInt(
		fmt.Sprintf("%.0f", claims[TokenKeyUserID]),
		10, 64)
	if err != nil {
		return nil, err
	}

	userID, err = svc.authRepo.Fetch(int(userID))
	if err != nil {
		return nil, err
	}

	accessToken, err := svc.generateAccessJwt(int(userID))
	if err != nil {
		return nil, err
	}

	claims, err = svc.RefreshTokenData(refreshToken)
	if err != nil {
		return nil, err
	}

	expire, err := strconv.ParseInt(
		fmt.Sprintf("%.0f", claims[TokenKeyExpire]),
		10, 64)
	if err != nil {
		return nil, err
	}

	return &dto.AccessTokenDto{
		Token:  accessToken,
		Expire: time.Unix(expire, 0),
	}, nil
}

func (svc *AuthService) AccessTokenData(accessToken string) (jwt.MapClaims, error) {
	jwtToken, err := svc.getToken(accessToken, configs.Env.AccessSecretKey)
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok && !jwtToken.Valid {
		return nil, fmt.Errorf("access token valid error")
	}

	return claims, nil
}

func (svc *AuthService) RefreshTokenData(refreshToken string) (jwt.MapClaims, error) {
	token, err := svc.getToken(refreshToken, configs.Env.RefreshSecretKey)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("Refresh token valid error")
	}

	return claims, nil
}

func (svc *AuthService) generateAccessJwt(userID int) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		TokenKeyAuthorized: true,
		TokenKeyUserID:     userID,
		TokenKeyExpire:     time.Now().Add(time.Minute * 60).Unix(),
	}).SignedString([]byte(configs.Env.AccessSecretKey))
}

func (svc *AuthService) generateRefreshJwt(userID int) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		TokenKeyUserID: userID,
		TokenKeyExpire: time.Now().Add(time.Hour * 24 * 7 * 2).Unix(),
	}).SignedString([]byte(configs.Env.RefreshSecretKey))
}

func (svc *AuthService) generateToken(userID int) (accessToken, refreshToken string, err error) {
	accessToken, err = svc.generateAccessJwt(userID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = svc.generateRefreshJwt(userID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (svc *AuthService) getToken(jwtString, key string) (*jwt.Token, error) {
	return jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(
				"unexpected signing method: %v", token.Header["alg"],
			)
		}

		return []byte(key), nil
	})
}
