package ingress

import (
	"database/sql"
	"djeurnie/api/internal/models"
	"fmt"
	"log"
)

type Service interface {
	All() *models.IngressList
	Get(ingressId string) (*models.Ingress, error)
}

type PlanetScaleDbService struct {
	svc *sql.DB
}

func NewPlanetScaleDbService(svc *sql.DB) Service {
	return &PlanetScaleDbService{
		svc: svc,
	}
}

func (s *PlanetScaleDbService) All() *models.IngressList {
	res, err := s.svc.Query("SELECT id, display_name FROM ingress")
	defer res.Close()

	ingressList := models.IngressList{}

	if err != nil {
		return &ingressList
	}

	for res.Next() {
		item := models.Ingress{}
		err := res.Scan(&item.Id, &item.DisplayName)
		if err != nil {
			log.Fatal(err)
			return &ingressList
		}
		ingressList.Items = append(ingressList.Items, item)
	}

	return &ingressList
}

func (s *PlanetScaleDbService) Get(ingressId string) (*models.Ingress, error) {
	item := models.Ingress{}
	err := s.svc.QueryRow("SELECT id, display_name FROM ingress WHERE id = ?", ingressId).Scan(&item.Id, &item.DisplayName)

	if err != nil {
		fmt.Println("error on GetItem: ", err)
		return nil, err
	}

	return &item, nil
}
