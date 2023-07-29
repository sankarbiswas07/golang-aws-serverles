package handlers

import(
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sankarbiswas07/golang-aws-serverles/pkg/user"
)

func CreateUser(){

}

func GetUser(){
	
}

func UpdateUser(){
	
}

func DeleteUser(){
	
}
func UnhandledMethod(){
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       "Method not allowed",
	}, nil
}