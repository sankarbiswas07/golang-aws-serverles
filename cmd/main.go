package main

import(
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sankarbiswas07/golang-aws-serverles/pkg/handlers"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	dynaClient *dynamodb.DynamoDB
)

func main()  {
	awsSession, err = session.NewSessionWithOptions(
		session.Options{
			Profile: "sankar.geotech.admin.practice"
		}
	)
	if err!= nil{
		return
	}
	dynaClient = dynamodb.New(awsSession)
  lambda.Start(lambdaHandler)
}

const tableName = "test-user"

func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
		case http.MethodPost:
			return handlers.CreateUser(ctx, request, tableName, dynaClient)

		case http.MethodGet:
			return handlers.GetUser(ctx, request, tableName, dynaClient)

		case http.MethodPut:
			return handlers.UpdateUser(ctx, request, tableName, dynaClient)

		case http.MethodDelete:
			return handlers.DeleteUser(ctx, request, tableName, dynaClient)

		default:
			return handlers.UnhandledMethod()
	}
}