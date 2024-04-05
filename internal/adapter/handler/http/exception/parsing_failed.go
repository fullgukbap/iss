package exception

import "github.com/gofiber/fiber/v2"

type ParsingFailedException struct {
	builtIn
}

func NewParsingFailedExceptino(err error) *ParsingFailedException {
	return &ParsingFailedException{
		builtIn: newBuiltIn("사용자의 요청을 해석해던 중 에러가 발생했습니다. 값을 잘 보냈는지 확인해주세요.", fiber.StatusBadRequest, err),
	}
}
