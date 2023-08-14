package ingress

import (
	"context"
	"djeurnie/api/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Service interface {
	All(tenant models.Tenant) *models.IngressList
}

type DynamoDbService struct {
	svc       *dynamodb.Client
	tableName string
}

func (s *DynamoDbService) All(tenant models.Tenant) *models.IngressList {
	out, err := s.svc.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(s.tableName),
		KeyConditionExpression: aws.String("TenantId = :TenantId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":TenantId": &types.AttributeValueMemberS{Value: tenant.Id},
		},
	})

	if err != nil {
		return &models.IngressList{}
	}

	if out.Count == 0 {
		return &models.IngressList{}
	}

	ingressList := models.IngressList{}

	for _, item := range out.Items {
		ingressList.Items = append(ingressList.Items, models.Ingress{
			Id:          item["IngressId"].(*types.AttributeValueMemberS).Value,
			DisplayName: item["DisplayName"].(*types.AttributeValueMemberS).Value,
		})
	}

	return &ingressList
}

func NewDynamoDbService(svc *dynamodb.Client, tableName string) Service {
	return &DynamoDbService{
		svc:       svc,
		tableName: tableName,
	}
}
