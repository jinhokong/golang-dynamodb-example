package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type RecommendHistoryItem struct {
	category  string `json:"category"`
	createdAt string `json:"createdAt"`
	ID        string `json:"ID"`
}

func GetItem(svc *dynamodb.DynamoDB) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("RecommendHistory"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String("01ba5e60-c1d8-4b21-b2dd-9039f806087e")},
		}})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	item := RecommendHistoryItem{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed", err))
	}
}
