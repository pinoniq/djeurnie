package handlers

import (
	"database/sql"
	"djeurnie/api/internal/models"
	"djeurnie/api/internal/service/egress"
	"djeurnie/api/internal/transport"
	"github.com/gofiber/fiber/v2"
)

type EgressHandler interface {
	List(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
}

type eHandler struct {
	Egress egress.Service
}

type egressDetailResponse struct {
	Id          string `json:"id" xml:"id"`
	DisplayName string `json:"displayName" xml:"displayName"`
}

type egressResponse struct {
	Id          string `json:"id" xml:"id"`
	DisplayName string `json:"displayName" xml:"displayName"`
}

type egressListResponse struct {
	Status string           `json:"status" xml:"status"`
	Egress []egressResponse `json:"items" xml:"items"`
}

func (h *eHandler) List(c *fiber.Ctx) error {
	listRes := egressListResponse{
		Status: "ok",
	}

	tenantsEgress := h.Egress.All()

	for _, item := range tenantsEgress.Items {
		listRes.Egress = append(listRes.Egress, egressResponse{
			Id:          item.Id,
			DisplayName: item.DisplayName,
		})
	}

	return transport.Encode(c, &listRes)
}

func (h *eHandler) Get(c *fiber.Ctx) error {
	egressId := c.Params("Id")

	item, err := h.Egress.Get(egressId)

	if err != nil {
		return err
	}

	return transport.Encode(c, &egressDetailResponse{
		Id:          item.Id,
		DisplayName: item.DisplayName,
	})
}

type EgressHandlerFactory struct {
	eHandler EgressHandler
	Egress   egress.Service
}

func (f EgressHandlerFactory) Supports(config *models.ApiRouteConfig) bool {
	return config.Target == "egress"
}

func (f EgressHandlerFactory) CreateHandler(config *models.ApiRouteConfig) fiber.Handler {
	e, err := f.Egress.Get(config.TargetId)

	if err != nil {
		return func(ctx *fiber.Ctx) error {
			return fiber.NewError(501, "Not Implemented")
		}
	}

	switch e.Config.Handler {
	case "list":
		return f.eHandler.List
	case "item":
		return f.eHandler.Get
	}

	return func(ctx *fiber.Ctx) error {
		return fiber.NewError(501, "Not Implemented")
	}
}

func NewEgressHandlerFactory(svc *sql.DB) HandlerFactory {
	s := egress.NewPlanetScaleDbService(svc)

	return EgressHandlerFactory{
		Egress: s,
		eHandler: &eHandler{
			Egress: s,
		},
	}
}
