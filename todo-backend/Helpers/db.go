package Helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var DB *mongo.Database

func init(){
	clientOptions := options.Client().ApplyURI("mongodb://MongoMaster:MongoPassword@vue-go-docker_database_1")
	client,err := mongo.Connect(context.TODO(),clientOptions)

	if err != nil{
		log.Fatal(err)
	}

	if err == nil{
		log.Println("Connection Succeeded!")
	}

	err = client.Ping(context.TODO(),nil)

	if err != nil{
		log.Fatal(err)
	}

	DB = client.Database("ToDos")
}
