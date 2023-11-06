package amqp

import (
	"encoding/json"
	"rabbit/lib/msgqueue"

	"github.com/streadway/amqp"
)

type amqpEventEmitter struct {
	connection *amqp.Connection
	exchange   string
	events     chan *emittedEvent
}

type emittedEvent struct {
	event msgqueue.Event
	error chan error
}

// AMQP 채널은 thread-safe 하지 않기 때문에, 이벤트를 보내기 위해 채널을 사용할 때는
// 이벤트를 보내기 전에 채널을 생성해야 한다.
// thread-safe하지 않다는 것은 AMQP 채널이 여러 스레드나 고루틴에 의해 동시에 안전하게 사용되도록 디자인되지 않았다는 것을 의미한다.
func NewAMQPEventEmitter(conn *amqp.Connection, exchange string) (msgqueue.EventEmitter, error) {
	emitter := &amqpEventEmitter{
		connection: conn,
		exchange:   exchange,
	}

	err := emitter.setup()
	if err != nil {
		return nil, err
	}

	return emitter, nil
}

func (a *amqpEventEmitter) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	//return channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
	return channel.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil)
}

func (a *amqpEventEmitter) Emit(event msgqueue.Event) error {
	ch, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	jsonDoc, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": event.EventName()},
		Body:        jsonDoc,
		ContentType: "application/json",
	}

	return ch.Publish(a.exchange, event.EventName(), false, false, msg)
	//return ch.Publish("events", event.EventName(), false, false, msg)
}
