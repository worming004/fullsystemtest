package main

import (
	"context"
	"log"
	"net/http"

	openapi "github.com/worming004/fullsystemtest/productservice/api-generated/go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	log.Printf("Server started")

	mongodb, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:aStrong!Password741@db:27017"))
	if err != nil {
		panic(err)
	}
	err = mongodb.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	mongodb.Ping(context.TODO(), readpref.Primary())
	DefaultApiService := newProductService(mongodb)
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
