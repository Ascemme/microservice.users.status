package handler

import (
	"github.com/Ascemme/microservice.users.status/pkg/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sssss"))
}
