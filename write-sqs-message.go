package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}
	svc := sqs.New(sess, &aws.Config{Region: aws.String("eu-west-1")})
	params := &sqs.SendMessageInput{
		MessageBody:  aws.String("a line or sequence of people or vehicles awaiting their turn to be attended to or to proceed."),
		QueueUrl:     aws.String("https://sqs.eu-west-1.amazonaws.com/259764441114/Learn"),
		DelaySeconds: aws.Int64(1),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"learn": {
				DataType:    aws.String("String"),
				StringValue: aws.String("queue"),
			},
		},
	}
	resp, err := svc.SendMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
}
