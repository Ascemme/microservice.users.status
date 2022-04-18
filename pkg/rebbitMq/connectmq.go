package rebbitMq

import (
	"errors"
	"github.com/streadway/amqp"
	"log"
)

type ChannelMQ struct {
	ChanMq       *amqp.Channel
	NameKey      string
	NameExchange string
}

func NewChannelMQ(chanMq *amqp.Channel, nameKey string, nameExchange string) *ChannelMQ {
	return &ChannelMQ{ChanMq: chanMq, NameKey: nameKey, NameExchange: nameExchange}
}

func ErrorGeter(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)

	}
}

func ConnectionMQ(nameKey, nameExchange string) (*amqp.Channel, error) {

	conn, err := amqp.Dial("amqp://ascemme:passwordqwe@localhost:5672/")
	if err != nil {
		return nil, errors.New("connection")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.New("chnael")
	}
	_, err = ch.QueueDeclare(nameKey, false, false, false, false, nil)

	ch.ExchangeDeclare(
		nameExchange, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	return ch, nil
}
