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

func del() {

	initDb()
	c := client.Database("go_db").Collection("Student")
	ctx := context.TODO()

	dr, err := c.DeleteMany(ctx, bson.D{{"age", 20}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dr.DeletedCount: %v\n", dr.DeletedCount)
}

func main() {
	del()
}
*/