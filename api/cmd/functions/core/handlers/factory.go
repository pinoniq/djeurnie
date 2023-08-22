package handlers

import (
	"database/sql"
	"djeurnie/api/internal/models"
	"github.com/gofiber/fiber/v2"
)

type HandlerFactory interface {
	Supports(config *models.ApiRouteConfig) bool
	CreateHandler(config *models.ApiRouteConfig) fiber.Handler
}

type Factory struct {
	svc      *sql.DB
	handlers []HandlerFactory
}

func (f *Factory) CreateHandler(route *models.ApiRoute) fiber.Handler {
	for _, handler := range f.handlers {
		if handler.Supports(&route.Config) {
			return handler.CreateHandler(&route.Config)
		}
	}

	return func(ctx *fiber.Ctx) error {
		return fiber.NewError(501, "Not Implemented")
	}
}

func NewFactory(svc *sql.DB) Factory {
	var routeHandlers []HandlerFactory

	routeHandlers = append(routeHandlers, NewEgressHandlerFactory(svc))
	routeHandlers = append(routeHandlers, NewIngressHandlerFactory(svc))

	return Factory{
		svc:      svc,
		handlers: routeHandlers,
	}
}
