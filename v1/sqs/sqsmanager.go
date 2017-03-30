package sqs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var svc *sqs.SQS
var qURL string

const queuename = "GO_API_QUEUE"

var messageCounter int

func NewSession() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a SQS service client.
	svc = sqs.New(sess)

	// create Queue
	resultCreateQueue, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(queuename),
		Attributes: map[string]*string{
			"DelaySeconds":                  aws.String("60"),
			"MessageRetentionPeriod":        aws.String("86400"),
			"ReceiveMessageWaitTimeSeconds": aws.String("5"),
		},
	})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	qURL = *resultCreateQueue.QueueUrl

	fmt.Println("Message queue successfully spun up", *resultCreateQueue.QueueUrl)

}
