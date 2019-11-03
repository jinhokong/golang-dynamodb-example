package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type InputType struct {
	Category            string   `json:"category"`
	NotPreferProductIDs []string `json:"notPreferProductIDs"`
	PreferProductIDs    []string `json:"preferProductIDs"`
	TagIDs              []string `json:"tagIDs"`
	UserID              string   `json:"userID"`
}

type RecommendHistoryItem struct {
	Category  string    `json:"category"`
	CreatedAt string    `json:"createdAt"`
	ID        string    `json:"ID"`
	Input     InputType `json:"input"`
}

func GetItem(svc *dynamodb.DynamoDB) []RecommendHistoryItem {
	Items := make([]RecommendHistoryItem, 0, 100)
	var lek map[string]*dynamodb.AttributeValue
	for {

		filt := expression.Name("category").Equal(expression.Value("CONCEALER"))

		expr, err := expression.NewBuilder().
			WithFilter(filt).
			Build()

		result, err := svc.Scan(&dynamodb.ScanInput{
			TableName:                 aws.String("RecommendHistory"),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			ExclusiveStartKey:         lek,
			FilterExpression:          expr.Filter(),
		})

		if result.LastEvaluatedKey == nil {
			break
		}

		lek = result.LastEvaluatedKey

		fmt.Println(lek)

		if err != nil {
			panic(fmt.Sprintf("Failed", err))
		}
		target := []RecommendHistoryItem{}

		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &target)
		if err != nil {
			panic(fmt.Sprintf("Failed", err))
		}
		Items = append(Items, target...)
	}
	return Items
}
