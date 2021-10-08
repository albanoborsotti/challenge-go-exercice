package services

import (
	"example.com/ml_tech/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func SavePersonDatabase(person *models.Person) error {

	db := GetDynamoDB()

	_, err := db.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(person.Id),
			},
			"dna_type": {
				S: aws.String(person.DnaType),
			},
			"dna": {
				SS: aws.StringSlice(person.Dna),
			},
		},
		TableName: aws.String("DNA"),
	})

	if err != nil {
		return err
	}
	return nil
}
