package http

import (
	"letsgo-mini-is/internal/adapter/handler/http/dto"
	"letsgo-mini-is/internal/adapter/handler/http/exception"

	"github.com/gofiber/fiber/v2"
)

func customErrorHandler(c *fiber.Ctx, err error) error {

	switch err.(type) {
	case *exception.InternalErrorException:
		e := err.(*exception.InternalErrorException)
		return c.Status(e.Status).JSON(&dto.GeneralResponse{
			Code:    e.Status,
			Message: e.Error(),
		})

	case *exception.ParsingFailedException:
		e := err.(*exception.ParsingFailedException)
		return c.Status(e.Status).JSON(&dto.GeneralResponse{
			Code:    e.Status,
			Message: e.Error(),
		})

	case *exception.InvalidUUIDException:
		e := err.(*exception.InvalidUUIDException)
		return c.Status(e.Status).JSON(&dto.GeneralResponse{
			Code:    e.Status,
			Message: e.Error(),
		})

	}

	return fiber.DefaultErrorHandler(c, err)
	// return c.SendStatus(fiber.StatusNotFound)
}
