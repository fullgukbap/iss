package exception

import "github.com/gofiber/fiber/v2"

type InvalidUUIDException struct {
	builtIn
}

func NewInvalidUUIDException(err error) *InvalidUUIDException {
	return &InvalidUUIDException{
		builtIn: newBuiltIn("uuid 값이 유효하지 않습니다. UUID값을 확인해주세요.", fiber.StatusBadRequest, err),
	}
}
