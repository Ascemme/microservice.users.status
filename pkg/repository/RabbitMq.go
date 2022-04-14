package repository

import (
	"context"
	"fmt"
	"github.com/Ascemme/microservice.users.status/pkg/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type RebbitMqRepo struct {
	*mongo.Collection
}

func (db *RebbitMqRepo) CreateStatus(user model.User) error {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	result, err := db.Collection.InsertOne(ctx, ser())
	if err != nil {
		fmt.Println("err1")
		fmt.Println(err)
	}

	sids, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		fmt.Println("ne ok")
	}
	fmt.Println(sids.Hex())
}

func (db *RebbitMqRepo) GetStatusById(id string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (db *RebbitMqRepo) GetStatusByUid(uid int) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (db *RebbitMqRepo) UpdateStatus(id string, user model.User) error {
	//TODO implement me
	panic("implement me")
}

func (db *RebbitMqRepo) DeleteSatus(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRebbitMqRepo(collection *mongo.Collection) *RebbitMqRepo {
	return &RebbitMqRepo{Collection: collection}
}
