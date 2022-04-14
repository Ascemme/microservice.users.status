package service

import (
	"encoding/json"
	"fmt"
	"github.com/Ascemme/microservice.users.status/pkg/model"
	"github.com/streadway/amqp"
)

func (s *Service) ServiceMq(ch <-chan amqp.Delivery) {
	for msgs := range ch {
		var msg model.Massage
		json.Unmarshal(msgs.Body, &msg)
		s.rabbiMQLogic(msg)
	}
}

func (s *Service) rabbiMQLogic(msg model.Massage) {
	switch msg.Value {
	case "dislikes":

	case "likse":

	case "follow":
	case "comments":
	case "subscribers":
	case "posts":
	case "pages":
	default:
		fmt.Println("tut")
		fmt.Println(msg)
	}
}

func (s *Service) gettingValue(uid int) (string, error) {
	id, err := s.repo.GetId(uid)
	if err != nil {
		return "", err
	}
	return id, err
}

func (s *Service) Dislike(msg model.Massage) error {
	id, err := s.gettingValue(msg.Uid)
	if err != nil || id == "" {
		return err
	}
	return s.
}
