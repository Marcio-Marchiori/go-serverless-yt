package user

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"encoding/json"
	"errors"
)

var(
	ErrorFailedToFetchRecord = "failed to fetch record"
)

type User struct{
	Email		string `json:"email"`
	FirstName	string `json:firstName`
	LastName	string `json:lastName`
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*User,error){
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email":{
				S: aws.String(email)
			}
		},
		TableName: aws.String(tableName)
	}

	result, err := dynaClient.GetItem(input)
	if err!=nil{
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
}

func FetchUsers()(){

}

func CreateUser()(){
	
}

func UpdateUser()(){
	
}

func DeleteUser() error{
	
}