package transport

import (
	"djeurnie/api/internal/models"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func encode[T any](c *fiber.Ctx, res T) error {
	accepts := c.Accepts("application/json", "application/x-protobuf", "application/xml")

	if accepts == "" {
		return errors.New("missing accept header")
	}

	switch accepts {
	case "application/json":
		return c.JSON(res)

	case "application/xml":
		return c.XML(res)
	}

	return errors.New("unsupported accept header")
}

func WrapEncodingWithTenant[T any](handler func(c *fiber.Ctx, t models.Tenant) (T, error)) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		rawTenant := c.Locals("tenant")

		if rawTenant == nil {
			return c.SendStatus(fiber.StatusNotAcceptable)
		}

		tenant := rawTenant.(models.Tenant)

		res, err := handler(c, tenant)
		if err != nil {
			return c.SendStatus(fiber.StatusNotAcceptable)
		}

		return encode(c, res)
	}
}