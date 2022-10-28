package types

import "github.com/gofiber/fiber/v2"

var (
	ErrServerInternal            = fiber.NewError(fiber.StatusInternalServerError, "잠시후 다시 이용하여 주시기 바랍니다")
	ErrPasswordNotCompareConfirm = fiber.NewError(fiber.StatusBadRequest, "비밀번호와 일치하지 않습니다")
	ErrPasswordInvalid           = fiber.NewError(fiber.StatusUnauthorized, "비밀번호가 일치하지 않습니다")
	ErrUserNotExist              = fiber.NewError(fiber.StatusBadRequest, "사용자가 존재하지 않습니다")
	ErrUserExist                 = fiber.NewError(fiber.StatusConflict, "사용자가 이미 존재합니다")
	ErrUnauthorized              = fiber.NewError(fiber.StatusConflict, "잘못된 인증입니다")
	ErrInvalidParameter          = fiber.NewError(fiber.StatusBadRequest, "잘못된 요청입니다")
	ErrPostNotExist              = fiber.NewError(fiber.StatusBadRequest, "존재하지 않는 게시물입니다.")
)
