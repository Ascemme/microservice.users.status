package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError22(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func connection() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		failOnError2(err, "Failed to connect to RabbitMQ")
		conn.Close()
	}
	return conn, err
}

func cannel(queueName string, connection *amqp.Connection) (<-chan amqp.Delivery, error) {
	ch, err := connection.Channel()
	if err != nil {
		failOnError2(err, "Failed to open a channel")
		ch.Close()
		return nil, err
	}
	q, err := ch.QueueDeclare(
		"golang-queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError2(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError2(err, "Failed to register a consumer")
	return msgs, nil

}

func main() {

	forever := make(chan bool)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
