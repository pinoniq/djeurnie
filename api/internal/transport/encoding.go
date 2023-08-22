package transport

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func Encode[T any](c *fiber.Ctx, res T) error {
	// @see middleware.ExtensionAsFormatter
	accepts := c.Locals(fiber.HeaderAccept).(string)

	switch accepts {
	case "application/json":
		return c.JSON(res)

	case "application/xml":
		return c.XML(res)
	}

	return errors.New("unsupported accept header")
}

func WrapEncoding[T any](handler func(c *fiber.Ctx) (T, error)) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		res, err := handler(c)
		if err != nil {
			return c.SendStatus(fiber.StatusNotAcceptable)
		}

		return Encode(c, res)
	}
}
