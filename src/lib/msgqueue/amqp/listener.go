package amqp

import (
	"encoding/json"
	"fmt"
	"rabbit/contracts"
	"rabbit/lib/msgqueue"

	"github.com/streadway/amqp"
	//amqp "github.com/rabbitmq/amqp091-go"
)

type amqpEventListener struct {
	connection *amqp.Connection
	queue      string
	exchange   string
	//mapper     msgqueue.EventMapper
}

func NewAMQPEventListner(conn *amqp.Connection, exchange string, queue string) (msgqueue.EventListener, error) {
	listener := &amqpEventListener{
		connection: conn,
		queue:      queue,
		exchange:   exchange,
		//mapper:     msgqueue.NewEventMapper(),
	}
	err := listener.setup()
	if err != nil {
		return nil, err
	}
	return listener, nil
}
func (a *amqpEventListener) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return nil
	}
	defer channel.Close()
	err = channel.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}
	// err 이미 사용했으므로 :=가 아닌 _를 사용해야 함
	_, err = channel.QueueDeclare(a.queue, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("could not declare queue %s: %s", a.queue, err)
	}
	return nil
}

// <-chan은 읽기 전용 채널을 의미한다.
func (a *amqpEventListener) Listen(eventNames ...string) (<-chan msgqueue.Event, <-chan error, error) {
	channel, err := a.connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	defer channel.Close()
	for _, eventName := range eventNames {
		//if err := channel.QueueBind(a.queue, eventName, "events", false, nil); err != nil {
		// Create binding between queue and exchange for each listened event type
		if err := channel.QueueBind(a.queue, eventName, a.exchange, false, nil); err != nil {
			//return nil, nil, err
			return nil, nil, fmt.Errorf("could not bind eventName %s to exchange %s: %s", eventName, a.queue, err)
		}
	}

	// Set up QoS options
	// err = channel.Qos(1, 0, false)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("could not configure QoS: %s", err)
	// }

	msgs, err := channel.Consume(a.queue, "", false, false, false, false, nil)
	if err != nil {
		return nil, nil, err
		//return nil, nil, fmt.Errorf("could not consume queue: %s", err)
	}
	events := make(chan msgqueue.Event, 9)
	errors := make(chan error)
	//the reason for the go routine here is that we want to return the events channels to the caller of the function immediately
	//and not wait for the for loop to finish before returning the channels

	go func() {
		//defer close(events)
		//defer close(errors)
		for msg := range msgs {
			rawEventName, ok := msg.Headers["x-event-name"]
			if !ok {
				errors <- fmt.Errorf("msg did not contain x-event-name header")
				msg.Nack(false, false)
				continue
			}
			eventName, ok := rawEventName.(string)
			if !ok {
				errors <- fmt.Errorf("x-event-name header is not a string, but %t", rawEventName)
				msg.Nack(false, false)
				continue
			}

			var event msgqueue.Event
			switch eventName {
			case "eventCreated":
				event = new(contracts.EventCreatedEvent)
			case "locationCreated":
				event = &contracts.LocationCreatedEvent{}
			case "eventBooked":
				event = &contracts.EventBookedEvent{}
			default:
				errors <- fmt.Errorf("=event type %s is unknown", eventName)
			}
			err := json.Unmarshal(msg.Body, event)
			if err != nil {
				errors <- err
				continue
			}
			events <- event
			msg.Ack(false)
		}
	}()
	return events, errors, nil

}

// <-chan은 읽기 전용 채널을 의미한다.
/*func (a *amqpEventListener) Listen(eventNames ...string) (chan msgqueue.Event, chan error, error) {
	restartListener := func() (chan msgqueue.Event, chan error, error) {
		// Reopen the channel if needed
		channel, err := a.connection.Channel()
		if err != nil {
			return nil, nil, err
		}
		defer channel.Close()

		// Rebind the queues
		for _, eventName := range eventNames {
			if err := channel.QueueBind(a.queue, eventName, a.exchange, false, nil); err != nil {
				return nil, nil, fmt.Errorf("could not bind eventName %s to exchange %s: %s", eventName, a.queue, err)
			}
		}

		msgs, err := channel.Consume(a.queue, "", false, false, false, false, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("could not consume queue: %s", err)
		}

		events := make(chan msgqueue.Event, 1)
		errors := make(chan error)

		go func() {
			for msg := range msgs {
				rawEventName, ok := msg.Headers["x-event-name"]
				if !ok {
					errors <- fmt.Errorf("msg did not contain x-event-name header")
					msg.Nack(false, false)
					continue
				}
				eventName, ok := rawEventName.(string)
				if !ok {
					errors <- fmt.Errorf("x-event-name header is not a string, but %t", rawEventName)
					msg.Nack(false, false)
					continue
				}

				var event msgqueue.Event
				switch eventName {
				case "eventCreated":
					event = new(contracts.EventCreatedEvent)
				case "locationCreated":
					event = &contracts.LocationCreatedEvent{}
				case "eventBooked":
					event = &contracts.EventBookedEvent{}
				default:
					errors <- fmt.Errorf("=event type %s is unknown", eventName)
				}
				err := json.Unmarshal(msg.Body, event)
				if err != nil {
					errors <- err
					continue
				}
				events <- event
				msg.Ack(false)
			}
		}()

		return events, errors, nil
	}

	events, errors, err := restartListener()
	if err != nil {
		return nil, nil, err
	}

	// Check for errors in the background and restart the listener if needed
	go func() {
		for {
			select {
			case err := <-errors:
				fmt.Printf("Error processing message: %s. Restarting listener...\n", err)
				close(events)
				close(errors)
				events, errors, _ = restartListener()
			}
		}
	}()

	return events, errors, nil
}*/
