package handlers

import (
	"djeurnie/api/internal/models"
	"github.com/gofiber/fiber/v2"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
	Tenant string `json:"tenant"`
}

func HealthCheck(c *fiber.Ctx, tenant models.Tenant) (*HealthCheckResponse, error) {
	return &HealthCheckResponse{
		Status: "ok",
		Tenant: tenant.Id,
	}, nil
}
