package handlers

import (
	"djeurnie/api/internal/models"
	"github.com/gofiber/fiber/v2"
)

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Tenant  string `json:"tenant"`
	Message string `json:"message"`
}

func HealthCheck(c *fiber.Ctx) (*HealthCheckResponse, error) {
	rawTenant := c.Locals("tenant")

	if rawTenant == nil {
		return &HealthCheckResponse{
			Status:  "error",
			Message: "No tenant found",
		}, nil
	}

	tenant := rawTenant.(models.Tenant)

	return &HealthCheckResponse{
		Status: "ok",
		Tenant: tenant.Id,
	}, nil
}
