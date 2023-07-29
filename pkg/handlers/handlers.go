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

var (
	ErrorMethodNotAllowed = "method not allowed"
)

type ErrorBody struct{
	ErrorMsg *string `json: "error, remove empty"`
}

func CreateUser(ctx context.Context, req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB)(*http.APIGatewayProxyResponse, error){

}

func GetUser(ctx context.Context, req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB)(*http.APIGatewayProxyResponse, error){
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err != nil{
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOk, result)
	}
	result, err := user.FetchUses(tableName, dynaClient)
	if err != nil{
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOk, result)
}

func UpdateUser(ctx context.Context, req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB)(*http.APIGatewayProxyResponse, error){
	
}

func DeleteUser(ctx context.Context, req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB)(*http.APIGatewayProxyResponse, error){
	
}
func UnhandledMethod()(*http.APIGatewayProxyResponse, error){
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       ErrorMethodNotAllowed,
	}, nil
}