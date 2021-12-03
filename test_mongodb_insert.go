/* package main

import (
	"context"
	"fmt"
	"log"

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

func insert() {

	initDb()

	s1 := Student{
		Name: "tom",
		Age:  20,
	}

	c := client.Database("go_db").Collection("Student")
	ior, err := c.InsertOne(context.TODO(), s1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("ior.InsertedID: %v\n", ior.InsertedID)
	}
	// defer collection
}

func insertMany() {
	initDb()

	c := client.Database("go_db").Collection("Student")

	s1 := Student{
		Name: "kiite",
		Age:  20,
	}

	s2 := Student{
		Name: "rose",
		Age:  20,
	}

	stus := []interface{}{s1, s2}

	imr, err := c.InsertMany(context.TODO(), stus)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("imr.InsertedIDs: %v\n", imr.InsertedIDs)

}

func main() {
	// insert()
	insertMany()

}
*/