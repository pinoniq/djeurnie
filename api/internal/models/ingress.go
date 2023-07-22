package models

type Ingress struct {
	TenantID string `dynamodbav:"tenantId" json:"tenantId" xml:"tenantId"`
	Id       string `dynamodbav:"ingressId" json:"id" xml:"id"`
}

type IngressList struct {
	Items []Ingress `json:"tenantId" xml:"tenantId"`
	Id    string    `json:"id" xml:"id"`
}
