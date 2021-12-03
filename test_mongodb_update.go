/* package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name string
	Age  int
}

var client *mongo.Client

func initDb() {
	co := options.Client().ApplyURI("mongodb://localhost:27017")
	c, err := mongo.Connect(context.TODO(), co)
	if err != nil {
		log.Fatal(err)
	}

	err2 := c.Ping(context.TODO(), nil)
	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println("连接成功！")
	}
	client = c
}

func update() {

	initDb()
	c := client.Database("go_db").Collection("Student")
	ctx := context.TODO()
	update := bson.D{{"$set", bson.D{{"name", "big kite"}, {"Age", 22}}}}
	ur, err := c.UpdateMany(ctx, bson.D{{"name", "kiite"}}, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.ModifiedCount: %v\n", ur.ModifiedCount)

}

func main() {
	update()
}
*/