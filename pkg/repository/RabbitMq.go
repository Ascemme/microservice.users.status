package repository

import (
	"context"
	"github.com/Ascemme/microservice.users.status/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type RebbitMqRepo struct {
	*mongo.Collection
}

func (db *RebbitMqRepo) CreateStatus(user model.User) error {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	result, err := db.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	_, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return err
	}
	return err
}

func (db *RebbitMqRepo) GetStatusById(id string) (string, error) {
	var myuser model.User
	ou, _ := primitive.ObjectIDFromHex(id)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	filter := bson.M{"_id": ou}
	result := db.Collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return id, result.Err()
	}
	if err := result.Decode(&myuser); err != nil {
		return "", err
	}

	return myuser.Id, nil
}

func (db *RebbitMqRepo) GetStatusByUid(uid int) (model.User, error) {
	var myuser model.User
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	filter := bson.M{"uid": uid}
	result := db.Collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return myuser, result.Err()
	}
	if err := result.Decode(&myuser); err != nil {
		return myuser, err
	}
	return myuser, nil
}

func (db *RebbitMqRepo) GetAll() ([]model.User, error) {
	var myusers []model.User
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	filter := bson.M{}
	options := options.Find()

	cur, err := db.Find(ctx, filter, options)
	if err != nil {
		return myusers, err
	}

	for cur.Next(ctx) {
		var myuser model.User
		err := cur.Decode(&myuser)
		if err != nil {
			log.Fatal(err)
		}
		myusers = append(myusers, myuser)
	}

	return myusers, nil
}

func (db *RebbitMqRepo) UpdateStatus(user model.User) error {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	objId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objId}
	user.Id = ""
	update := bson.D{{"$set", user}}
	_, err = db.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return err
}

func (db *RebbitMqRepo) DeleteStatus(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	se, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": se}
	_, err = db.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return err
}
func NewRebbitMqRepo(collection *mongo.Collection) *RebbitMqRepo {
	return &RebbitMqRepo{Collection: collection}
}
