package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

var defaultAccepts = "application/json"

func DeciderOfEncodings() fiber.Handler {

	return func(c *fiber.Ctx) error {
		currentPath := c.Path()

		if strings.HasSuffix(currentPath, ".json") {
			c.Locals(fiber.HeaderAccept, "application/json")
			c.Path(strings.TrimSuffix(currentPath, ".json"))
			return c.Next()
		}

		if strings.HasSuffix(currentPath, ".xml") {
			c.Locals(fiber.HeaderAccept, "application/xml")
			c.Path(strings.TrimSuffix(currentPath, ".xml"))
			return c.Next()
		}

		accepts := c.Accepts("application/json", "application/x-protobuf", "application/xml")

		if accepts == "" {
			accepts = defaultAccepts
		}

		c.Locals(fiber.HeaderAccept, accepts)

		return c.Next()
	}
}
