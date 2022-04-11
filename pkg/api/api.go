package api

import (
	"github.com/Ascemme/microservice.users.status/pkg/handler"
	"github.com/gorilla/mux"
)

type Api struct {
	handler *handler.Handler
}

func NewApi(handler *handler.Handler) *Api {
	return &Api{handler: handler}
}

func (a *Api) ApiGroup(r *mux.Router) {
	r.HandleFunc("/10", a.handler.GetStatus).Methods("GET")
}
