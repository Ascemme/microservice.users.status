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
	defer client.Disconnect(ctx)

	newDatabase := client.Database("Status")

	newCollection := NewRepository(newDatabase.Collection("Users"))
	//newCollection.TestCol()
	//newCollection.FindeOne("62554d9a6b3ceac027cd2eac")
	//newCollection.GetAll()
	//newCollection.DeleteColomn("62554d9a6b3ceac027cd2eac")
	//newCollection.UpdetePut("62554d933558add9ee6275c6")
	newCollection.GetId(14)
	return newDatabase.Collection("Users")
}
