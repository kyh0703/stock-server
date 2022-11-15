package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/kyh0703/stock-server/config"
	"github.com/kyh0703/stock-server/database"
	"golang.org/x/crypto/bcrypt"
)

type TokenMetaData struct {
	AccessToken         string
	AccessUUID          string
	AccessTokenExpires  int64
	RefreshUUID         string
	RefreshToken        string
	RefreshTokenExpires int64
}

type AuthService struct {
	authRepo AuthRepository
}

func (svc *AuthService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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

func (svc *AuthService) Login(id int) (*TokenMetaData, error) {
	// create token
	token, err := svc.generateToken(id)
	if err != nil {
		return nil, err
	}

	// save the redis
	if err := svc.saveToken(id, token); err != nil {
		return nil, err
	}

	return token, nil
}

func (svc *AuthService) Logout(jwtString string) error {
	uuid, err := svc.GetUUIDByAccessToken(jwtString)
	if err != nil {
		return err
	}

	deleted, err := svc.authRepo.Delete(uuid)
	if err != nil || deleted == 0 {
		return errors.New("failed to remove token")
	}

	return nil
}

func (svc *AuthService) Refresh(jwtString string) (*TokenMetaData, error) {
	uuid, err := svc.getUUIDByRefreshToken(jwtString)
	if err != nil {
		return nil, err
	}

	userID, err := svc.getUserIDByRefreshToken(jwtString)
	if err != nil {
		return nil, err
	}

	deleted, err := svc.authRepo.Delete(uuid)
	if err != nil || deleted == 0 {
		return nil, errors.New("failed to remove token")
	}

	token, err := svc.Login(userID)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (svc *AuthService) FindUserIDByUUID(UUID string) (int, error) {
	id, err := database.Redis.Get(UUID).Int()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (svc *AuthService) GetAccessToken(jwtString string) (jwt.MapClaims, error) {
	jwtToken, err := svc.getToken(jwtString, config.Env.AccessSecretKey)
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok && !jwtToken.Valid {
		return nil, errors.New("Access token valid error")
	}

	return claims, nil
}

func (svc *AuthService) GetRefreshToken(jwtString string) (jwt.MapClaims, error) {
	token, err := svc.getToken(jwtString, config.Env.RefreshSecretKey)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("Refresh token valid error")
	}

	return claims, nil
}

func (svc *AuthService) GetUUIDByAccessToken(jwtString string) (string, error) {
	jwtToken, err := svc.getToken(jwtString, config.Env.AccessSecretKey)
	if err != nil {
		return "", err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return "", errors.New("failed extract token metadata")
	}

	accessUUID, ok := claims["access_uuid"].(string)
	if !ok {
		return "", errors.New("failed to get access uuid")
	}

	return accessUUID, nil
}

func (svc *AuthService) generateAccessToken(id int, UUID string, expire int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["access_uuid"] = UUID
	claims["user_id"] = id
	claims["exp"] = expire

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env.AccessSecretKey))
}

func (svc *AuthService) generateRefreshToken(id int, UUID string, expire int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["refresh_uuid"] = UUID
	claims["user_id"] = id
	claims["exp"] = expire

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env.RefreshSecretKey))
}

func (svc *AuthService) generateToken(userID int) (*TokenMetaData, error) {
	// Create access token metadata with expire 30min
	metadata := new(TokenMetaData)
	metadata.AccessTokenExpires = time.Now().Add(time.Minute * 30).Unix()
	metadata.AccessUUID = uuid.NewString()
	accessToken, err := svc.generateAccessToken(
		userID,
		metadata.AccessUUID,
		metadata.AccessTokenExpires,
	)
	if err != nil {
		return nil, err
	}
	metadata.AccessToken = accessToken

	// Create refresh token metadata with expire 7day
	metadata.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	metadata.RefreshUUID = uuid.NewString()
	refreshToken, err := svc.generateRefreshToken(
		userID,
		metadata.RefreshUUID,
		metadata.RefreshTokenExpires,
	)
	if err != nil {
		return nil, err
	}
	metadata.RefreshToken = refreshToken
	return metadata, nil
}

func (svc *AuthService) saveToken(userID int, token *TokenMetaData) error {
	// convert Unis to UTC
	var (
		accessTokenExpire  = time.Unix(token.AccessTokenExpires, 0)
		refreshTokenExpire = time.Unix(token.RefreshTokenExpires, 0)
		now                = time.Now()
	)
	if err := svc.authRepo.InsertToken(
		userID,
		token.AccessUUID,
		accessTokenExpire,
		now,
	); err != nil {
		return err
	}
	if err := svc.authRepo.InsertToken(
		userID,
		token.RefreshUUID,
		refreshTokenExpire,
		now,
	); err != nil {
		return err
	}
	return nil
}

func (svc *AuthService) getToken(jwtString, key string) (*jwt.Token, error) {
	return jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
}

func (service *AuthService) getUUIDByRefreshToken(jwtString string) (string, error) {
	jwtToken, err := service.getToken(jwtString, config.Env.RefreshSecretKey)
	if err != nil {
		return "", err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return "", errors.New("failed extract token metadata")
	}

	refreshUUID, ok := claims["refresh_uuid"].(string)
	if !ok {
		return "", errors.New("failed to get refresh uuid")
	}

	return refreshUUID, nil
}

func (service *AuthService) getUserIDByRefreshToken(jwtString string) (int, error) {
	jwtToken, err := service.getToken(jwtString, config.Env.RefreshSecretKey)
	if err != nil {
		return 0, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return 0, errors.New("failed extract token metadata")
	}

	userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 64)
	if err != nil {
		return 0, err
	}

	return int(userID), nil
}
