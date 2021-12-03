/* package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func find() {
	ctx := context.TODO()
	defer client.Disconnect(ctx)
	c := client.Database("go_db").Collection("Student")

	c2, err := c.Find(ctx, bson.D{{"name", "tom"}})
	if err != nil {
		log.Fatal(err)
	}
	defer c2.Close(ctx)

	for c2.Next(ctx) {
		var result bson.D
		c2.Decode(&result)

		fmt.Printf("result: %v\n", result)
		fmt.Printf("result: %v\n", result.Map())

	}

}

func main() {
	initDb()
	find()
}
 */