package sqs

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/shiva0705/goApi/v1/models"
)

func Push(feedback models.Feedback) {

	feedbackjson, err := json.Marshal(feedback)

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"VideoId": &sqs.MessageAttributeValue{
				DataType:    aws.String("Number"),
				StringValue: aws.String(strconv.Itoa(feedback.VideoId)),
			},
			"Like": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(strconv.FormatBool(feedback.Like)),
			},
		},
		MessageBody: aws.String(string(feedbackjson)),
		QueueUrl:    &qURL,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Feedback message pushed", *result.MessageId)

	time.Sleep(5 * time.Second)
	pull()
}
