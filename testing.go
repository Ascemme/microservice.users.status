package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Sera struct {
	Id  int
	Bas string
	Vas string
}

func failOnError2(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func main() {
	result := make(chan string)
	ch, _, err := declareQueue2("test1", false, false, false, false, nil)
	fmt.Println(err)
	//createEx1(ch, "testTutut")
	geter(result, ch)
	s := <-result
	fmt.Println(s)
}

func declareQueue2(name string, durable, autoDelete,
	exclusive, noWait bool, argp amqp.Table) (*amqp.Channel, *amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://ascemme:passwordqwe@localhost:5672/")
	if err != nil {
		return nil, nil, errors.New("connection")
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, errors.New("chnael")
	}
	_, err = ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, argp)
	return ch, conn, nil
}

func createEx1(ch *amqp.Channel, name string) {
	ch.ExchangeDeclare(
		name,     // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

}

func geter(result chan string, ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		"golang-queue1", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	failOnError2(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,     // queue
		"testCons", // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args

	)
	err = ch.QueueBind(
		q.Name,          // queue name
		"golang-queue1", // routing key
		"testTutut",     // exchange
		false,
		nil)
	failOnError2(err, "Failed to register a consumer")
	ses := Sera{}
	for i := range msgs {
		json.Unmarshal([]byte(i.Body), &ses)
		fmt.Println(ses.Vas)
	}

}
