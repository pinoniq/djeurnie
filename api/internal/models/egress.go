package models

import (
	"fmt"
)

type EgressDataBlob struct {
	headers []string
	body    []string
}

type Egress struct {
	TenantID                 string                         `dynamodbav:"tenantId"`
	Id                       string                         `dynamodbav:"id"`
	IdentificationStrategies []IdentificationStrategyConfig `dynamodbav:"identificationStrategies"`
}

type IdentificationStrategyConfig struct {
	Id       string
	Strategy string
	Config   []string
}

func (isc *IdentificationStrategyConfig) createStrategy() (IdentificationStrategy, error) {
	switch isc.Strategy {
	case "PropertyMatcher":
		return PropertyMatcherStrategy{
			Config: isc.Config,
		}, nil
	}

	return nil, fmt.Errorf("createStrategy: No strategy found for %s", isc.Strategy)
}

type IdentifiedDataModel struct {
	Id                     string
	IdentificationStrategy IdentificationStrategy
	Egress                 Egress
}

type IdentificationStrategy interface {
	Identify(input EgressDataBlob) IdentifiedDataModel
}

type PropertyMatcherStrategy struct {
	Config []string
}

func (pms PropertyMatcherStrategy) Identify(input EgressDataBlob) IdentifiedDataModel {
	return IdentifiedDataModel{
		Id:                     "test",
		IdentificationStrategy: pms,
	}
}
