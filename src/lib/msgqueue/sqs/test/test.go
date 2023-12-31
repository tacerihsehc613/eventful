package main

import (
	"fmt"
	"rabbit/contracts"
	"rabbit/lib/msgqueue/sqs"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	queueName := "eventqueue"     // Replace with your SQS queue name
	awsRegion := "ap-northeast-2" // Replace with your AWS region

	//sess, err := session.NewSession()
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		fmt.Println("Error creating AWS session:", err)
		return
	}

	emitter, err := sqs.NewSQSEventEmitter(sess, queueName)
	if err != nil {
		fmt.Println("Error creating SQS emitter:", err)
		return
	}

	// Sample event to emit
	event := contracts.EventCreatedEvent{
		ID:         "123",
		Name:       "Sample Event",
		LocationID: "ERICA",
		Start:      time.Now(),
		End:        time.Now().Add(time.Hour),
	}

	// Emit the event
	err = emitter.Emit(&event)
	if err != nil {
		fmt.Println("Error emitting event:", err)
		return
	}

	fmt.Println("Event emitted successfully!")

	// Create an SQS event listener
	listener, err := sqs.NewSQSListener(sess, queueName, 10, 5, 30)
	if err != nil {
		fmt.Println("Error creating SQS listener:", err)
		return
	}

	// Listen for events
	eventCh, errorCh, err := listener.Listen("eventCreated")
	if err != nil {
		fmt.Println("Error setting up listener:", err)
		return
	}

	// Consume events
	go func() {
		for {
			select {
			case receivedEvent := <-eventCh:
				fmt.Printf("Received event: %+v\n", receivedEvent)
			case err := <-errorCh:
				fmt.Println("Error receiving event:", err)
			}
		}
	}()

	// Sleep to allow time for events to be processed
	time.Sleep(time.Second * 10)

	// Note: In a real-world scenario, you would have a listener/consumer
	// that processes events from the SQS queue.
	// You can create a separate program or goroutine for that.
}
