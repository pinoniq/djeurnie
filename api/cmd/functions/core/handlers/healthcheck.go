package handlers

import (
	"djeurnie/api/internal/transport"
	"github.com/gofiber/fiber/v2"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheck(c *fiber.Ctx) error {
	return transport.Encode(c, &HealthCheckResponse{
		Status: "ok",
	})
}
