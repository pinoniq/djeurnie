package egress

import (
	"database/sql"
	"djeurnie/api/internal/models"
	"fmt"
	"log"
)

type Service interface {
	All() *models.EgressList
	Get(egressId string) (*models.Egress, error)
}

type PlanetScaleDbService struct {
	svc *sql.DB
}

func NewPlanetScaleDbService(svc *sql.DB) Service {
	return &PlanetScaleDbService{
		svc: svc,
	}
}

func (s *PlanetScaleDbService) All() *models.EgressList {
	res, err := s.svc.Query("SELECT id, display_name FROM egress")
	defer res.Close()

	egressList := models.EgressList{}

	if err != nil {
		return &egressList
	}

	for res.Next() {
		item := models.Egress{}
		err := res.Scan(&item.Id, &item.DisplayName)
		if err != nil {
			log.Fatal(err)
			return &egressList
		}
		egressList.Items = append(egressList.Items, item)
	}

	return &egressList
}

func (s *PlanetScaleDbService) Get(egressId string) (*models.Egress, error) {
	item := models.Egress{}
	err := s.svc.QueryRow("SELECT id, display_name, config FROM egress WHERE id = ?", egressId).Scan(&item.Id, &item.DisplayName, &item.Config)

	if err != nil {
		fmt.Println("error on GetItem: ", err)
		return nil, err
	}

	return &item, nil
}
