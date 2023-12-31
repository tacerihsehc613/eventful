package sqs

import (
	"errors"
	"rabbit/lib/msgqueue"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSListener struct {
	mapper              msgqueue.EventMapper
	sqsSvc              *sqs.SQS
	queueURL            *string
	maxNumberOfMessages int64
	waitTime            int64
	visibilityTimeout   int64
}

func NewSQSListener(s *session.Session, queueName string, maxMsgs, wtTime, visTO int64) (listener msgqueue.EventListener, err error) {
	if s == nil {
		s, err = session.NewSession()
		if err != nil {
			return
		}
	}
	svc := sqs.New(s)
	QUResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		return
	}
	listener = &SQSListener{
		sqsSvc:              svc,
		queueURL:            QUResult.QueueUrl,
		mapper:              msgqueue.NewEventMapper(),
		maxNumberOfMessages: maxMsgs,
		waitTime:            wtTime,
		visibilityTimeout:   visTO,
	}
	return
}

func (sqsListener *SQSListener) Listen(events ...string) (<-chan msgqueue.Event, <-chan error, error) {
	if sqsListener == nil {
		return nil, nil, errors.New("SQSListener: the Listen() method was called on a nil pointer")
	}
	eventCh := make(chan msgqueue.Event)
	errorCh := make(chan error)
	go func() {
		for {
			sqsListener.receiveMessage(eventCh, errorCh)
		}
	}()
	return eventCh, errorCh, nil
}

func (sqsListener *SQSListener) receiveMessage(eventCh chan msgqueue.Event, errorCh chan error, events ...string) {
	recvMsgResult, err := sqsListener.sqsSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            sqsListener.queueURL,
		MaxNumberOfMessages: aws.Int64(sqsListener.maxNumberOfMessages),
		WaitTimeSeconds:     aws.Int64(sqsListener.waitTime),
		VisibilityTimeout:   aws.Int64(sqsListener.visibilityTimeout),
	})
	if err != nil {
		errorCh <- err
		return
	}
	bContinue := false
	for _, msg := range recvMsgResult.Messages {
		value, ok := msg.MessageAttributes["EventType"]
		if !ok {
			continue
		}
		eventName := aws.StringValue(value.StringValue)
		for _, event := range events {
			if strings.EqualFold(eventName, event) {
				bContinue = true
				break
			}
		}
		if !bContinue {
			continue
		}

		message := aws.StringValue(msg.Body)
		event, err := sqsListener.mapper.MapEvent(eventName, []byte(message))
		if err != nil {
			errorCh <- err
			return
		}
		eventCh <- event
		_, err = sqsListener.sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      sqsListener.queueURL,
			ReceiptHandle: msg.ReceiptHandle,
		})
		if err != nil {
			errorCh <- err
		}

	}

}

func (sqsListener *SQSListener) Mapper() msgqueue.EventMapper {
	return sqsListener.mapper
}
