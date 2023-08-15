package ingress

import (
	"context"
	"database/sql"
	"djeurnie/api/internal/models"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

type Service interface {
	All(tenant models.Tenant) *models.IngressList
	Get(tenant models.Tenant, ingressId string) (*models.Ingress, error)
}

type PlanetScaleDbService struct {
	svc *sql.DB
}

func NewPlanetScaleDbService(svc *sql.DB) Service {
	return &PlanetScaleDbService{
		svc: svc,
	}
}

func (s *PlanetScaleDbService) All(tenant models.Tenant) *models.IngressList {
	res, err := s.svc.Query("SELECT id, tenant_id, display_name FROM ingress WHERE tenant_id = ?", tenant.Id)
	defer res.Close()

	ingressList := models.IngressList{}

	if err != nil {
		return &ingressList
	}

	for res.Next() {
		item := models.Ingress{}
		err := res.Scan(&item.Id, &item.TenantID, &item.DisplayName)
		if err != nil {
			log.Fatal(err)
			return &ingressList
		}
		ingressList.Items = append(ingressList.Items, item)
	}

	return &ingressList
}

func (s *PlanetScaleDbService) Get(tenant models.Tenant, ingressId string) (*models.Ingress, error) {
	item := models.Ingress{}
	err := s.svc.QueryRow("SELECT id, tenant_id, display_name FROM ingress WHERE tenant_id = ? AND ingress_id = ?", tenant.Id, ingressId).Scan(&item.Id, &item.TenantID, &item.DisplayName)

	if err != nil {
		fmt.Println("error on GetItem: ", err)
		return nil, err
	}

	return &item, nil
}

type DynamoDbService struct {
	svc       *dynamodb.Client
	tableName string
}

func NewDynamoDbService(svc *dynamodb.Client, tableName string) Service {
	return &DynamoDbService{
		svc:       svc,
		tableName: tableName,
	}
}

func (s *DynamoDbService) All(tenant models.Tenant) *models.IngressList {
	out, err := s.svc.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(s.tableName),
		KeyConditionExpression: aws.String("TenantId = :TenantId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":TenantId": &types.AttributeValueMemberS{Value: tenant.Id},
		},
	})

	ingressList := models.IngressList{}

	if err != nil {
		return &ingressList
	}

	if out.Count == 0 {
		return &ingressList
	}

	for _, item := range out.Items {
		ingressList.Items = append(ingressList.Items, models.Ingress{
			Id:          item["IngressId"].(*types.AttributeValueMemberS).Value,
			DisplayName: item["DisplayName"].(*types.AttributeValueMemberS).Value,
		})
	}

	return &ingressList
}

func (s *DynamoDbService) Get(tenant models.Tenant, ingressId string) (*models.Ingress, error) {
	out, err := s.svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("TableName"),
		Key: map[string]types.AttributeValue{
			"TenantId":  &types.AttributeValueMemberS{Value: tenant.Id},
			"IngressId": &types.AttributeValueMemberS{Value: ingressId},
		},
	})

	if err != nil {
		fmt.Println("error on GetItem: ", err)
		return nil, err
	}

	if out.Item == nil {
		return nil, errors.New("ingress not found")
	}

	item := models.Ingress{}

	err = attributevalue.UnmarshalMap(out.Item, &item)

	if err != nil {
		fmt.Println("error on unmarshal: ", err)
		return nil, err
	}

	return &item, nil
}
