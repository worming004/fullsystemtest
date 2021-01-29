package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	openapi "github.com/worming004/fullsystemtest/productservice/api-generated/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var internalError = openapi.ImplResponse{
	Code: http.StatusInternalServerError,
	Body: "internal error",
}

type productService struct {
	collection *mongo.Collection
}

func newProductService(client *mongo.Client) *productService {
	return &productService{
		collection: client.Database("productsdb").Collection("products"),
	}
}

func (ps productService) ProductsGet(ctx context.Context) (openapi.ImplResponse, error) {
	cursor, err := ps.collection.Find(ctx, bson.D{})
	if err != nil {
		return internalError, err
	}

	var result []openapi.Product
	err = cursor.All(ctx, &result)
	if err != nil {
		return internalError, err
	}

	return openapi.Response(http.StatusOK, result), nil
}

func (ps productService) ProductsIdDelete(_ context.Context, _ int32) (openapi.ImplResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (ps productService) ProductsIdGet(_ context.Context, _ int32) (openapi.ImplResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (ps productService) ProductsPost(ctx context.Context, p openapi.Product) (openapi.ImplResponse, error) {
	if p.Id != "" {
		return openapi.Response(http.StatusBadRequest, "do not set id"), nil
	}
	id, _ := uuid.NewUUID()
	p.Id = id.String()
	_, err := ps.collection.InsertOne(ctx, p)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}
