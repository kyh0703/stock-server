package types

import "github.com/gofiber/fiber/v2"

var (
	ErrUserNotExist = fiber.NewError(fiber.StatusBadRequest, "사용자가 존재하지 않습니다")
	ErrUserExist    = fiber.NewError(fiber.StatusConflict, "사용자가 이미 존재합니다")
	ErrAuthFail     = fiber.NewError(fiber.StatusConflict, "인증이 실패하였습니다")
)
