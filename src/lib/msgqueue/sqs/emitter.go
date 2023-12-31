package sqs

import (
	"encoding/json"
	"rabbit/lib/msgqueue"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSEmitter struct {
	sqsSvc   *sqs.SQS
	QueueURL *string
}

/* func NewSQSEventEmitter(s *session.Session, queueName string) (msgqueue.EventEmitter, error) {
	if s == nil {
		var err error
		s, err = session.NewSession()
		if err != nil {
			return nil, err
		}
	}
	svc := sqs.New(s)
	QUResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		return nil, err
	}
	emitter := &SQSEmitter{
		sqsSvc:   svc,
		QueueURL: QUResult.QueueUrl,
	}
	return emitter, nil
} */

func NewSQSEventEmitter(s *session.Session, queueName string) (emitter msgqueue.EventEmitter, err error) {
	if s == nil {
		s, err = session.NewSession()
		if err != nil {
			return
		}
	}
	svc := sqs.New(s)
	QUResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		return
	}
	emitter = &SQSEmitter{
		sqsSvc:   svc,
		QueueURL: QUResult.QueueUrl,
	}
	return
}

func (sqsEmit *SQSEmitter) Emit(event msgqueue.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	_, err = sqsEmit.sqsSvc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"event_name": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(event.EventName()),
			},
		},
		MessageBody: aws.String(string(data)),
		QueueUrl:    sqsEmit.QueueURL,
	})
	return err
}
