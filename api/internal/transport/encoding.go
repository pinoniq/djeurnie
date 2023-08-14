package transport

import (
	"djeurnie/api/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func encode[T any](c *fiber.Ctx, res T, accepts string) error {

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

		// @see middleware.ExtensionAsFormatter
		accepts := c.Locals(fiber.HeaderAccept).(string)

		go logResponse(res, accepts)

		return encode(c, res, accepts)
	}
}

func logResponse[T any](res T, accepts string) {
	time.Sleep(2 * time.Second)
	// Do stuff that continues doing stuff after the response has been sent
	fmt.Println("response:")
	switch accepts {
	case "application/json":
		rawResponse, err := json.Marshal(res)
		if err == nil {
			fmt.Println(string(rawResponse))
		} else {
			fmt.Println(err)
		}
	}
}
