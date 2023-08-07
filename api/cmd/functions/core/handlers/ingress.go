package handlers

import (
	"context"
	"djeurnie/api/internal/database"
	"djeurnie/api/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gofiber/fiber/v2"
)

type IngressResponse struct {
	Id          string `json:"id" xml:"id"`
	DisplayName string `json:"displayName" xml:"displayName"`
}

type IngressListResponse struct {
	Status  string            `json:"status" xml:"status"`
	Tenant  string            `json:"tenant" xml:"tenant"`
	Ingress []IngressResponse `json:"items" xml:"items"`
}

func IngressList(c *fiber.Ctx, tenant models.Tenant) (*IngressListResponse, error) {
	listRes := IngressListResponse{
		Status: "ok",
		Tenant: tenant.Id,
	}

	tableName := "ingress"
	svc := database.GetDynamodbSession()

	out, err := svc.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("TenantId = :TenantId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":TenantId": &types.AttributeValueMemberS{Value: tenant.Id},
		},
	})

	if err != nil {
		return nil, err
	}

	for _, item := range out.Items {
		listRes.Ingress = append(listRes.Ingress, IngressResponse{
			Id:          item["IngressId"].(*types.AttributeValueMemberS).Value,
			DisplayName: item["DisplayName"].(*types.AttributeValueMemberS).Value,
		})
	}

	return &listRes, nil
}
