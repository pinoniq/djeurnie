package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Ingress struct {
	Id          string        `db:"id"`
	DisplayName string        `db:"display_name"`
	Config      IngressConfig `db:"config"`
}

type IngressConfig struct {
	Handler string `json:"handler"`
}

type IngressList struct {
	Items []Ingress
}

func (ic *IngressConfig) Value() (driver.Value, error) {
	return json.Marshal(ic)
}

func (ic *IngressConfig) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &ic)
}
