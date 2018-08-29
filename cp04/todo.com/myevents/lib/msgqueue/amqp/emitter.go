package amqp

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type amqpEventEmitter struct {
	connection *amqp.Connection
}

func (a *amqpEventEmitter) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
}

func NewAMQPEventEmitter(conn *amqp.Connection) (EventEmitter, error) {
	emitter := &amqpEventEmitter{
		connection: conn,
	}
	err := emitter.setup()
	if err != nil {
		return nil, err
	}
	return emitter, nil
}

func (a *amqpEventEmitter) Emit(event Event) error {
	jsonDoc, err := json.Marshal(event)
	if err != nil {
		return err
	}

	chann, err := a.connection.Channel()
	if err != nil {
		return err
	}
	defer chann.Close()

	msg := amqp.Publishing{
		Headers: amqpTable{"X-event-name": event.EventName()},
		Body: jsonDoc,
		ContentType: "application/json",
	}

	return chann.Publish(
		"events",
		events.EventName()
	)
}
