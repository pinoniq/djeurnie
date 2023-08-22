package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Egress struct {
	Id          string       `db:"id"`
	DisplayName string       `db:"display_name"`
	Config      EgressConfig `db:"config"`
}

type EgressConfig struct {
	Handler string `json:"handler"`
}

type EgressList struct {
	Items []Egress
}

func (ec *EgressConfig) Value() (driver.Value, error) {
	return json.Marshal(ec)
}

func (ec *EgressConfig) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &ec)
}
