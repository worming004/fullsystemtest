package main

import (
	"log"
	"net/http"

	openapi "github.com/worming004/fullsystemtest/productservice/api-generated/go"
)

func main() {
	log.Printf("Server started")

	DefaultApiService := newProductService()
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
