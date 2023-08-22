package apiRoutes

import (
	"database/sql"
	"djeurnie/api/internal/models"
)

type Service interface {
	AllForResourceGroup(resourceGroupId string) *models.ApiRouteList
}

type PlanetScaleDbService struct {
	svc *sql.DB
}

func NewPlanetScaleDbService(svc *sql.DB) Service {
	return &PlanetScaleDbService{
		svc: svc,
	}
}

func (s *PlanetScaleDbService) AllForResourceGroup(resourceGroupId string) *models.ApiRouteList {
	res, err := s.svc.Query("SELECT id, path, method, config FROM api_routes WHERE resource_group_id = ?", resourceGroupId)
	defer res.Close()

	apiRouteList := models.ApiRouteList{}

	if err != nil {
		return &apiRouteList
	}

	for res.Next() {
		item := models.ApiRoute{}
		err := res.Scan(&item.Id, &item.Path, &item.Method, &item.Config)
		if err != nil {
			return &apiRouteList
		}
		apiRouteList.Items = append(apiRouteList.Items, item)
	}

	return &apiRouteList
}
