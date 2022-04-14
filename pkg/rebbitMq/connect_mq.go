package rebbitMq

import (
	"errors"
	"github.com/streadway/amqp"
	"log"
)

type ChannelMQ struct {
	ChanMq *amqp.Channel
}

func NewChannelMQ(chanMq *amqp.Channel) *ChannelMQ {
	return &ChannelMQ{ChanMq: chanMq}
}

func ErrorGeter(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)

	}
}

func ConnectionMQ() (*amqp.Channel, error) {
	var name string = "test1"
	conn, err := amqp.Dial("amqp://ascemme:passwordqwe@localhost:5672/")
	if err != nil {
		return nil, errors.New("connection")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.New("chnael")
	}
	_, err = ch.QueueDeclare(name, false, false, false, false, nil)
	return ch, nil
}
