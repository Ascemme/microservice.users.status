package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type ste struct {
	iwer chan int
}

type Sera2 struct {
	Id  int    `json:"id"`
	Bas string `json:"bas"`
	Vas string `json:"vas"`
}

func (s *ste) res() {
	i := 80
	for {
		s.iwer <- i
		i++
	}

}
func main() {
	result := make(chan int)
	sss := ste{
		result,
	}
	ch, _, err := declareQueue("test1", false, false, false, false, nil)
	fmt.Println(err)
	createEx(ch, "testTutut")
	go sss.sender(ch)
	go sss.res()
	i := 0
	for {
		result <- i
		i++
	}

}

func declareQueue(name string, durable, autoDelete,
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

func createEx(ch *amqp.Channel, name string) {
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

func (s *ste) sender(ch *amqp.Channel) {
	result := s.iwer
	q, err := ch.QueueDeclare(
		"golang-queue1", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	fmt.Println(err, "Failed to declare a queue")

	// We set the payload for the message.
	var i int
	var ss Sera2
	for se := range result {
		body := fmt.Sprintf("Golang is awesome - Keep Moving Forward! %d", i)
		ss.Id = i
		ss.Vas = body
		ss.Vas = fmt.Sprintf("vas   Golang is awesome - Keep Moving Forward! %d", se)

		se, err := json.Marshal(ss)
		fmt.Println(err)

		fmt.Println(se)
		err = ch.Publish(
			"testTutut", // exchange
			q.Name,      // routing key
			false,       // mandatory
			false,       // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(se),
			})
		fmt.Println(err, "Failed to publish a message")
		log.Printf(" [x] Congrats, sending message: %s", body)
		time.Sleep(time.Second * 3)
		i++
	}

	// If there is an error publishing the message, a log will be displayed in the terminal.

}
