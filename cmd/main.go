package main

import (
	"fmt"
	server "github.com/Ascemme/microservice.users.status"
	"github.com/Ascemme/microservice.users.status/pkg/api"
	"github.com/Ascemme/microservice.users.status/pkg/handler"
	"github.com/Ascemme/microservice.users.status/pkg/rebbitMq"
	"github.com/Ascemme/microservice.users.status/pkg/repository"
	"github.com/Ascemme/microservice.users.status/pkg/service"
	"github.com/gorilla/mux"
)

func main() {
	// router
	r := mux.NewRouter()

	// references
	repositories := repository.NewRepository(repository.ConnectionDb())
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	apis := api.NewApi(handlers)

	// rabbitMq

	mqChan, err := rebbitMq.ConnectionMQ("Status", "ServiceOne")
	if err != nil {
		fmt.Println(err)
	}

	chanls := rebbitMq.NewChannelMQ(mqChan, "Status", "ServiceOne")
	cn := chanls.GetMassage()

	go services.ServiceMq(cn)

	//running server
	apis.ApiGroup(r)
	server.Run(r, repositories)
}
