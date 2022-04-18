package repository

import (
	"github.com/Ascemme/microservice.users.status/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type RabbitMQRepo interface {
	CreateStatus(user model.User) error
	GetStatusById(id string) (string, error)
	GetStatusByUid(uid int) (model.User, error)
	GetAll() ([]model.User, error)
	UpdateStatus(user model.User) error
	DeleteStatus(id string) error
}

//	*mongo.Collection
type Repository struct {
	RabbitMQRepo
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{
		NewRebbitMqRepo(collection),
	}
}
