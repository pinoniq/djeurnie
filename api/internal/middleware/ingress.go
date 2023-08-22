package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Config struct {
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter func(*fiber.Ctx) bool
}

func logger(res *fiber.Response) {
	time.Sleep(2 * time.Second)
	fmt.Println("logger")
	fmt.Println(res.StatusCode())
	fmt.Println(res.Body())
}

func NewIngress(config ...Config) func(*fiber.Ctx) error {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}

	return func(c *fiber.Ctx) error {
		// Filter request to skip middleware
		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}

		defer func() {
			go logger(c.Response())
		}()

		return c.Next()
	}
}
