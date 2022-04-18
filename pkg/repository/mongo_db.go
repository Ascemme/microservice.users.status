package repository

import (
	"context"
	"fmt"
	"github.com/Ascemme/microservice.users.status/pkg/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectionDb() *mongo.Collection {
	s := settings.NewSettings()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", s.LoginDB, s.PasswordDB, s.Host, s.Port)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	newDatabase := client.Database("Status")
	return newDatabase.Collection("Users")
}
