package user

import(
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var(
	ErrorFailedToFetchRecord = "failed to fetch record"
	ErrorFailedToUnmarshalRecord = "failed to unmarshal Dynamodb Scan record"
)

type User struct{
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func FetchUser(email string, tableName string, dynaClient *dynamodb.DynamoDB)(*User, error){
	input := &dynamodb.GetItemInput{
		Key: map[string] *dynamodb.AttributeValue{
			"email": { 
				S: aws.String(email)
			}
		},
		TableName: aws.String(tableName)
	}

	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new(User)
	err := dynamodbattribute.Unmarshal(page.Items, item)
	if err != nil {
				return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	return item, nil
}

func FetchUsers(tableName string, dynaClient *dynamodb.DynamoDB)([]User, error){
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName)
	}

	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]User)
	err := dynamodbattribute.UnmarshalList(page.Items, item)
	if err != nil {
				return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	return item, nil
}

func CreateUser()(){
	
}

func UpdateUser(){
	
}

func DeleteUser() error{
	// return fmt.Errorf("user.go: DeleteUser(): %w", err)
}
