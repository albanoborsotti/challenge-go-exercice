package services

import (
	"example.com/ml_tech/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetStatsService() models.DnaStats {
	humans, mutants := GetCount()
	var stat = models.DnaStats{
		CountMutantDna: mutants,
		CountHumanDna:  humans,
		Ratio:          float64(mutants) / float64(humans),
	}
	return stat
}

func GetCount() (int64, int64) {
	db := GetDynamoDB()

	mutantsData, _ := db.Query(&dynamodb.QueryInput{
		TableName: aws.String("DNA"),
		IndexName: aws.String("Dna-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			"dna_type": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("MUTANT"),
					},
				},
			},
		},
	})

	humansData, _ := db.Query(&dynamodb.QueryInput{
		TableName: aws.String("DNA"),
		IndexName: aws.String("Dna-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			"dna_type": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("HUMAN"),
					},
				},
			},
		},
	})

	return *humansData.Count, *mutantsData.Count

}
