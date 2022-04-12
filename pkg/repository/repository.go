package repository

import (
	"context"
	"fmt"
	"github.com/Ascemme/microservice.users.status/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Repository struct {
	*mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{Collection: collection}
}

func ser() model.User {
	ser := model.User{}
	ser.Uid = 14
	ser.Following = append(ser.Following, 4)
	ser.Page = append(ser.Page, model.Page{Id: 10})
	return ser
}

func ser1() model.User {
	ser := model.User{}
	ser.Uid = 14
	ser.Following = append(ser.Following, 4)
	ser.Page = append(ser.Page, model.Page{Id: 10})
	return ser
}
func ser2() model.User {
	ser := model.User{}
	ser.Uid = 15
	ser.Following = append(ser.Following, 4)
	ser.Page = append(ser.Page, model.Page{Id: 10})
	return ser
}
func ser3() model.User {
	ser := model.User{}
	ser.Uid = 14
	ser.Following = append(ser.Following, 4)
	post := model.Post{Id: 18, Comments: []int{2, 3, 4}}
	var pit []model.Post
	pit = append(pit, post)
	ser.Page = append(ser.Page, model.Page{Id: 10, Post: pit})
	return ser
}

func (db Repository) TestCol() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	result, err := db.Collection.InsertOne(ctx, ser3())
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
func (db Repository) FindeOne(str string) {
	var myuser model.User
	ou, _ := primitive.ObjectIDFromHex(str)
	fmt.Println(ou)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	filter := bson.M{"_id": ou}
	result := db.Collection.FindOne(ctx, filter)
	if result.Err() != nil {
		fmt.Println("Resalt eror!@#!!!")
	}
	if err := result.Decode(&myuser); err != nil {
		fmt.Println("ne dicodits !0981904")
		fmt.Println(err)
	}
	fmt.Println(myuser)
}

func (db Repository) GetAll() {

	var myuser1 []model.User
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	filter := bson.M{}
	options := options.Find()

	cur, err := db.Find(ctx, filter, options)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var myuser model.User
		err := cur.Decode(&myuser)
		if err != nil {
			log.Fatal(err)
		}

		myuser1 = append(myuser1, myuser)
	}
	fmt.Println(len(myuser1))
	fmt.Println(myuser1)
}

func (db Repository) UpdetePut(s string) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	se, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	filter := bson.D{{"_id", se}}

	var gotovoe bson.M
	userBytes, err := bson.Marshal(ser2())
	if err := bson.Unmarshal(userBytes, &gotovoe); err != nil {
		fmt.Println("gotovoe")
		fmt.Println(err)
	}
	update := bson.M{"$set": gotovoe}

	result, err := db.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.MatchedCount, result.ModifiedCount)
}

func (db Repository) DeleteColomn(s string) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	se, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	filter := bson.M{"_id": se}
	one, err := db.Collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(one.DeletedCount)
}
