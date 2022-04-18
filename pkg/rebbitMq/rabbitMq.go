package rebbitMq

import "github.com/streadway/amqp"

func (ch *ChannelMQ) GetMassage() <-chan amqp.Delivery {
	q, err := ch.ChanMq.QueueDeclare(
		ch.NameKey, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	ErrorGeter(err, "Failed to declare a queue")

	msgs, err := ch.ChanMq.Consume(
		q.Name,          // queue
		ch.NameExchange, // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args

	)
	ErrorGeter(err, "chan getting")
	err = ch.ChanMq.QueueBind(
		q.Name,          // queue name
		ch.NameKey,      // routing key
		ch.NameExchange, // exchange
		false,
		nil)
	ErrorGeter(err, "Failed to register a consumer")

	return msgs

}
