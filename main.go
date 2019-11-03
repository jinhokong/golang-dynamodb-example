package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-2"),
		Credentials: credentials.NewSharedCredentials("", "fitsme.dev"),
	})
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(sess)
	svc := dynamodb.New(sess)
	result := GetItem(svc)
	// // 출력파일 생성
	// fo, err := os.Create("./data.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer fo.Close()
	// _, err = fo.Write(result)
	file, _ := json.MarshalIndent(result, "", " ")

	_ = ioutil.WriteFile("recommend.json", file, 0644)

}
