package handlers

import(
	"context"
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

func GetUser(ctx context.Context, req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB)(*events.APIGatewayProxyResponse, error){
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err != nil{
			return apiResponse(
				http.StatusBadRequest,
				 ErrorBody{
					aws.String(err.Error()),
				},
			)
		}
		return apiResponse(
			http.StatusOK,
			 result,
		)
	}
	result, err := user.FetchUsers(tableName, dynaClient)
	if err != nil{
		return apiResponse(
			http.StatusBadRequest,
			 ErrorBody{
				aws.String(err.Error()),
			},
		)
	}
	return apiResponse(
		http.StatusOK,
		result,
	)
}


func CreateUser(ctx context.Context, req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB)(*events.APIGatewayProxyResponse, error){
	result, err := user.CreateUser(req, tableName, dynaClient)
		if err != nil{
			return apiResponse(
				http.StatusBadRequest,
				 ErrorBody{
					aws.String(err.Error()),
				},
			)
		}
		return apiResponse(
			http.StatusOK,
			 result,
		)
}


func UpdateUser(ctx context.Context, req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB)(*events.APIGatewayProxyResponse, error){
	result, err := user.UpdateUser(req, tableName, dynaClient)
	if err != nil{
		return apiResponse(
			http.StatusBadRequest,
			 ErrorBody{
				aws.String(err.Error()),
			},
		)
	}
	return apiResponse(
		http.StatusOK,
		result,
	)
}

func DeleteUser(ctx context.Context, req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB) (*events.APIGatewayProxyResponse, error) {
	err := user.DeleteUser(req, tableName, dynaClient)

	if err!= nil {
		return apiResponse(http.StatusBadRequest, 
			ErrorBody{
				aws.String(err.Error()),
			},
		)
	}
	return apiResponse(http.StatusOK, nil)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error){
	return apiResponse(
		http.StatusMethodNotAllowed,
		ErrorMethodNotAllowed,
	)
}