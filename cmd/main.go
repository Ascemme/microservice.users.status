package main

import (
	server "github.com/Ascemme/microservice.users.status"
	"github.com/Ascemme/microservice.users.status/pkg/api"
	"github.com/Ascemme/microservice.users.status/pkg/handler"
	"github.com/Ascemme/microservice.users.status/pkg/repository"
	"github.com/Ascemme/microservice.users.status/pkg/service"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	repositories := repository.NewRepository(repository.ConnectionDb())
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	apis := api.NewApi(handlers)

	apis.ApiGroup(r)
	server.Run(r)
}
