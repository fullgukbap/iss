package http

import (
	"errors"
	"letsgo-mini-is/internal/adapter/handler/http/exception"

	"github.com/gofiber/fiber/v2"
)

func customErrorHandler(c *fiber.Ctx, err error) error {
	if errors.As(err, &exception.ParsingFailedException{}) {
		exception := err.(*exception.ParsingFailedException)
		return c.Status(exception.Status).SendString(err.Error())
	}

	if errors.As(err, &exception.InternalErrorException{}) {
		exception := err.(*exception.InternalErrorException)
		return c.Status(exception.Status).SendString(err.Error())
	}

	if errors.As(err, &exception.InvalidUUIDException{}) {
		exception := err.(*exception.InvalidUUIDException)
		return c.Status(exception.Status).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNotFound)
}
