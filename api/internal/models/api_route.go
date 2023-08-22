package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ApiRouteList struct {
	Items []ApiRoute
}

type ApiRoute struct {
	Id     string         `db:"id"`
	Path   string         `db:"path"`
	Method string         `db:"method"`
	Config ApiRouteConfig `db:"config"`
}

type ApiRouteConfig struct {
	Target   string `json:"target"`
	TargetId string `json:"targetId"`
}

func (arc *ApiRouteConfig) Value() (driver.Value, error) {
	return json.Marshal(arc)
}

func (arc *ApiRouteConfig) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &arc)
}
