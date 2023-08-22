package handlers

import (
	"database/sql"
	"djeurnie/api/internal/models"
	"djeurnie/api/internal/service/ingress"
	"djeurnie/api/internal/transport"
	"github.com/gofiber/fiber/v2"
)

type IngressHandler interface {
	List(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
}

type handler struct {
	Ingress ingress.Service
}

type ingressDetailResponse struct {
	Id          string `json:"id" xml:"id"`
	DisplayName string `json:"displayName" xml:"displayName"`
}

type ingressResponse struct {
	Id          string `json:"id" xml:"id"`
	DisplayName string `json:"displayName" xml:"displayName"`
}

type ingressListResponse struct {
	Status  string            `json:"status" xml:"status"`
	Ingress []ingressResponse `json:"items" xml:"items"`
}

func (h *handler) List(c *fiber.Ctx) error {
	listRes := ingressListResponse{
		Status: "ok",
	}

	tenantsIngress := h.Ingress.All()

	for _, item := range tenantsIngress.Items {
		listRes.Ingress = append(listRes.Ingress, ingressResponse{
			Id:          item.Id,
			DisplayName: item.DisplayName,
		})
	}

	return transport.Encode(c, &listRes)
}

func (h *handler) Get(c *fiber.Ctx) error {
	ingressId := c.Params("Id")

	item, err := h.Ingress.Get(ingressId)

	if err != nil {
		return err
	}

	return transport.Encode(c, &ingressDetailResponse{
		Id:          item.Id,
		DisplayName: item.DisplayName,
	})
}

type IngressHandlerFactory struct {
	handler IngressHandler
}

func (f IngressHandlerFactory) Supports(config *models.ApiRouteConfig) bool {
	return config.Target == "ingress"
}

func (f IngressHandlerFactory) CreateHandler(config *models.ApiRouteConfig) fiber.Handler {
	switch config.TargetId {
	case "list":
		return f.handler.List
	case "get":
		return f.handler.Get
	}

	return f.handler.List
}

func NewIngressHandlerFactory(svc *sql.DB) HandlerFactory {
	return IngressHandlerFactory{
		handler: &handler{
			Ingress: ingress.NewPlanetScaleDbService(svc),
		},
	}
}
