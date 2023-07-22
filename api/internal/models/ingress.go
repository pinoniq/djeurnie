package models

type Ingress struct {
	TenantID string `dynamodbav:"tenantId"`
	Id       string `dynamodbav:"ingressId"`
}

type IngressList struct {
	Items []Ingress
}
