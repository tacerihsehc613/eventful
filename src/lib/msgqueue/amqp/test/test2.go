package main

import (
	"log"
	"rabbit/contracts"
	msgqueue_amqp "rabbit/lib/msgqueue/amqp"
	"time"

	"github.com/streadway/amqp"
	//amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// RabbitMQ connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Exchange name and queue name
	exchange := "testExchange"
	queue := "testQueue"

	// Event emitter setup
	emitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn, exchange)
	if err != nil {
		log.Fatalf("Failed to create event emitter: %v", err)
	}

	// Event listener setup
	listener, err := msgqueue_amqp.NewAMQPEventListner(conn, exchange, queue)
	if err != nil {
		log.Fatalf("Failed to create event listener: %v", err)
	}

	// Listen for the "eventCreated" event
	receivedEvents, errors, err := listener.Listen("eventCreated")
	if err != nil {
		log.Fatalf("Failed to start listening: %v", err)
	}

	// Wait for and print received events in a goroutine
	done := make(chan struct{}) // Channel to signal completion
	defer close(done)
	go func() {
		for {
			select {
			case receivedEvent := <-receivedEvents:
				switch e := receivedEvent.(type) {
				case *contracts.EventCreatedEvent:
					log.Printf("Received event22: %+v\n", e)
				default:
					log.Printf("Unknown event type: %T", e)
				}
			case err := <-errors:
				log.Printf("Received error: %v", err)
			}
		}
	}()
	//<-done

	// Emit an "EventCreatedEvent" message
	event := &contracts.EventCreatedEvent{
		ID:         "1234",
		Name:       "My Festa",
		LocationID: "456",
		Start:      time.Now(),
		End:        time.Now().Add(time.Hour),
	}

	err = emitter.Emit(event)
	if err != nil {
		log.Fatalf("Failed to emit event: %v", err)
	}
	<-done

	// Wait for and print received events
	/*for {
		select {
		case receivedEvent := <-receivedEvents:
			//fmt.Printf("Received event: %+v\n", receivedEvent)
			switch e := receivedEvent.(type) {
			case *contracts.EventCreatedEvent:
				log.Printf("Received event22: %+v\n", e)
			default:
				log.Printf("Unknown event type: %T", e)
			}
		case err := <-errors:
			log.Printf("Received error: %v", err)
		}
	}*/
}
