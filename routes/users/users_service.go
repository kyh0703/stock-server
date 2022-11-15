package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/routes/auth"
	"github.com/kyh0703/stock-server/routes/users/dto"
	"github.com/kyh0703/stock-server/types"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	usersRepo UsersRepository
	authSvc   auth.AuthService
}

func (svc *UsersService) Register(c *fiber.Ctx, req *dto.UsersRegisterRequest) error {
	if req.Password != req.PasswordConfirm {
		return c.App().ErrorHandler(c, types.ErrPasswordNotCompareConfirm)
	}

	hash, err := svc.HashPassword(req.Password)
	if err != nil {
		return err
	}

	exist, err := svc.usersRepo.IsExistByEmail(c.Context(), req.Email)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrUserNotExist)
	} else if exist {
		return c.App().ErrorHandler(c, types.ErrUserExist)
	}

	if _, err = svc.usersRepo.Insert(
		c.Context(),
		req.Username,
		req.Email,
		hash,
	); err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (svc *UsersService) Login(c *fiber.Ctx, req *dto.UsersLoginRequest) error {
	user, err := svc.usersRepo.FetchByEmail(c.Context(), req.Email)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrUserExist)
	}

	ok, err := svc.CompareHashPassword(user.Password, req.Password)
	if err != nil || !ok {
		return c.App().ErrorHandler(c, types.ErrPasswordInvalid)
	}

	token, err := svc.authSvc.Login(user.ID)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	res := &dto.UsersLoginResponse{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return c.JSON(res)
}

func (svc *UsersService) Logout(c *fiber.Ctx, token string) error {
	if err := svc.authSvc.Logout(token); err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (svc *UsersService) GetUserDetail(c *fiber.Ctx, req *dto.UsersProfileRequest) error {
	user, err := svc.usersRepo.FetchOne(c.Context(), req.ID)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrUserNotExist)
	}

	res := &dto.UsersProfileResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
	return c.JSON(res)
}

func (svc *UsersService) RefreshToken(c *fiber.Ctx, req *dto.UsersRefreshRequest) error {
	token, err := svc.authSvc.Refresh(req.RefreshToken)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token.RefreshToken
	cookie.HTTPOnly = true
	cookie.Secure = true
	c.Cookie(cookie)

	res := &dto.UsersRefreshResponse{
		AccessToken: token.AccessToken,
	}
	return c.JSON(res)
}

func (svc *UsersService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash: %w", err)
	}
	return string(hash), nil
}

func (svc *UsersService) CompareHashPassword(hash, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		if err != bcrypt.ErrMismatchedHashAndPassword {
			return false, fmt.Errorf("mismatch password: %w", err)
		}
		return false, err
	}
	return true, nil
}
