package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/internal/routes/users/dtos"
	"github.com/kyh0703/stock-server/internal/types"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	usersRepo   *UsersRepository
	authService *AuthService
}

func NewUsersService(usersRepo *UsersRepository, authService *AuthService) *UsersService {
	return &UsersService{
		usersRepo:   usersRepo,
		authService: authService,
	}
}

func (svc *UsersService) Register(c *fiber.Ctx, req *dtos.UsersRegisterRequest) error {
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

	if _, err = svc.usersRepo.Save(
		c.Context(),
		req.Username,
		req.Email,
		hash,
	); err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (svc *UsersService) Login(c *fiber.Ctx, req *dtos.UsersLoginRequest) error {
	user, err := svc.usersRepo.FetchByEmail(c.Context(), req.Email)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrUserExist)
	}

	ok, err := svc.CompareHashPassword(user.Password, req.Password)
	if err != nil || !ok {
		return c.App().ErrorHandler(c, types.ErrPasswordInvalid)
	}

	tokenData, err := svc.authService.Login(user.ID)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	// NOTE [Access Token 처리]
	// native web일경우에는 쿠키를 사용할 수 없으니
	// Header나 Body에 담아 전달하나 web 개발이기에 cookie에 넣어둠
	// 미들웨어 처리
	// 1. Access Token, Refresh Token 두개 다 발급
	//   - AccessToken Cookie에 저장
	//   - RefreshToken Redis에 저장
	// 2. AccessToken 만료, Refresh Token 만료 >> 회원가입
	// 3. AccessToken 만료, Refresh Token 유효 >> AccessToken 재발급
	// 4. AccessToken 유효, Refresh Token 만료 >> RefreshToken 재발급
	// 5. AccessToken 유효, Refresh Token 유효 >> next()
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenData.Token
	cookie.HTTPOnly = true
	cookie.Secure = true
	c.Cookie(cookie)

	res := &dtos.UsersLoginResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(res)
}

func (svc *UsersService) Logout(c *fiber.Ctx, token string) error {
	return c.SendStatus(fiber.StatusNoContent)
}

func (svc *UsersService) GetUserDetail(c *fiber.Ctx, req *dtos.UsersProfileRequest) error {
	user, err := svc.usersRepo.FetchOne(c.Context(), req.ID)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrUserNotExist)
	}

	res := &dtos.UsersProfileResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(res)
}

func (svc *UsersService) RefreshToken(c *fiber.Ctx, req *dtos.UsersRefreshRequest) error {
	tokenData, err := svc.authService.Refresh(req.RefreshToken)
	if err != nil {
		return c.App().ErrorHandler(c, types.ErrServerInternal)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenData.Token
	cookie.HTTPOnly = true
	cookie.Secure = true
	c.Cookie(cookie)

	return c.SendStatus(fiber.StatusOK)
}

func (svc *UsersService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash: %w", err)
	}

	return string(hash), nil
}

func (svc *UsersService) CompareHashPassword(hash, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	); err != nil {
		if err != bcrypt.ErrMismatchedHashAndPassword {
			return false, fmt.Errorf("mismatch password: %w", err)
		}
		return false, err
	}

	return true, nil
}
