package exception

import "github.com/gofiber/fiber/v2"

type InternalErrorException struct {
	builtIn
}

func NewInternalErrorException(err error, reasons ...string) *InternalErrorException {
	reason := "서버의 내부 오류가 발생했습니다. "
	if len(reasons) > 0 {
		reason = reasons[0]
	}
	return &InternalErrorException{
		builtIn: newBuiltIn(reason, fiber.StatusInternalServerError, err),
	}

}
