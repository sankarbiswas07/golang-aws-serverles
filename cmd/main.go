package main

import(
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sankarbiswas07/golang-aws-serverles/pkg"

	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
	dynaClient := dynamodb.New(awsSession)
  lambda.Start(lambdaHandler)
}

const tableName = "test-user"

func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
		case http.MethodPost:
			return handlers.postUser(ctx, request, tableName, dynaClient)

		case http.MethodGet:
			return handlers.getUser(ctx, request, tableName, dynaClient)

		case http.MethodPut:
			return handlers.putUser(ctx, request, tableName, dynaClient)

		case http.MethodDelete:
			return handlers.deleteUser(ctx, request, tableName, dynaClient)

		default:
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusMethodNotAllowed,
				Headers:    map[string]string{"Content-Type": "text/plain"},
				Body:       "Method not allowed",
			}, nil
	}
}