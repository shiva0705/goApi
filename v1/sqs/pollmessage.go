package sqs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/shiva0705/goApi/v1/data"
	"github.com/shiva0705/goApi/v1/models"
)

func startpolling() {
	pull()
}

func pull() {

	fmt.Println("Polling message")

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &qURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(0),
		WaitTimeSeconds:     aws.Int64(0),
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	if len(result.Messages) == 0 {
		fmt.Println("Received no messages")
	} else {

		var feedback models.Feedback
		var messageAttr = result.Messages[0].Attributes

		var id = messageAttr["VideoId"]
		var like = messageAttr["Like"]

		fmt.Printf("Message values: %v, %v", &id, &like)

		feedback = models.Feedback{VideoId: 1, Like: true}
		process(feedback)
		delete(result)
	}
}

func process(feedback models.Feedback) {

	var db = data.DbHandle()
	defer db.Close()

	data.UpdateFeedback(db, feedback)
}

func delete(result *sqs.ReceiveMessageOutput) {
	resultDelete, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &qURL,
		ReceiptHandle: result.Messages[0].ReceiptHandle,
	})

	if err != nil {
		fmt.Println("Delete Error", err)
		return
	}

	fmt.Println("Message Deleted", resultDelete)
}
