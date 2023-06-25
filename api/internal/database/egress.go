package database

import (
	"context"
	"djeurnie/api/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetEgressFromDynamoDb(TenantId string, Id string) (models.Egress, error) {
	svc := GetDynamodbSession()
	rawEgressData, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("Egress"),
		Key: map[string]types.AttributeValue{
			"tenantId": &types.AttributeValueMemberS{Value: TenantId},
			"id":       &types.AttributeValueMemberS{Value: Id},
		},
	})

	if nil != err {
		return models.Egress{}, err
	}

	var egress models.Egress
	err = attributevalue.UnmarshalMap(rawEgressData.Item, &egress)

	return egress, err
}
