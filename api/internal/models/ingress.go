package models

type Ingress struct {
	TenantID    string `dynamodbav:"TenantId"`
	Id          string `dynamodbav:"IngressId"`
	DisplayName string `dynamodbav:"DisplayName"`
}

type IngressList struct {
	Items []Ingress
}
