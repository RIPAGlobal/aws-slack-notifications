package message

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

type SlackMessageIdentifier struct {
	BuildID string `json:"buildID"`
	TS      string `json:"ts"`
}

func SlackTSLookup(buildID string) string {
	fmt.Printf("\n%s %s", "Looking up:", buildID)

	// Create the session that the DynamoDB service will use.
	awsSession := session.Must(session.NewSession())

	// Create the DynamoDB service client to make the query request with.
	svc := dynamodb.New(awsSession)

	mid := SlackMessageIdentifier{}

	buildLookupResult, err := svc.GetItem(
		&dynamodb.GetItemInput{
			AttributesToGet:          nil,
			ConsistentRead:           aws.Bool(true),
			ExpressionAttributeNames: nil,
			Key: map[string]*dynamodb.AttributeValue{
				"buildID": {S: aws.String(buildID)},
			},
			ProjectionExpression:   nil,
			ReturnConsumedCapacity: nil,
			TableName:              aws.String(os.Getenv("DYNAMO_TABLE_NAME")),
		})
	HandleTSErrors(err, buildID)

	err = dynamodbattribute.UnmarshalMap(buildLookupResult.Item, &mid)
	HandleTSErrors(err, buildID)

	return mid.TS
}

func SaveNewMessageTS(buildID string, slackTS string) {
	// Create the session that the DynamoDB service will use.
	awsSession := session.Must(session.NewSession())

	// Create the DynamoDB service client to make the query request with.
	svc := dynamodb.New(awsSession)

	mid := SlackMessageIdentifier{
		BuildID: buildID,
		TS:      slackTS,
	}

	newSlackMessageID, err := dynamodbattribute.MarshalMap(mid)
	HandleTSErrors(err, buildID)

	input := &dynamodb.PutItemInput{
		Item:      newSlackMessageID,
		TableName: aws.String(os.Getenv("DYNAMO_TABLE_NAME")),
	}

	_, err = svc.PutItem(input)
	HandleTSErrors(err, buildID)
}

// Lookup Error printer if exists - not much else we can do with them.
func HandleTSErrors(err error, buildID string) {
	if err != nil {
		fmt.Println("# Error:")
		fmt.Println(err.Error())
		fmt.Println("# Build ID:")
		fmt.Printf("\n%s", buildID)
	}
}
