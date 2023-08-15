package models

type Ingress struct {
	Id          string `dynamodbav:"IngressId" db:"id"`
	TenantID    string `dynamodbav:"TenantId" db:"tenant_id"`
	DisplayName string `dynamodbav:"DisplayName" db:"display_name"`
}

type IngressList struct {
	Items []Ingress
}
