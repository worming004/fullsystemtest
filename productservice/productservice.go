package main

import (
	"context"
	"net/http"

	openapi "github.com/worming004/fullsystemtest/productservice/api-generated/go"
	"go.mongodb.org/mongo-driver/mongo"
)

type productSlice []openapi.Product

type productService struct {
	collection *mongo.Collection
}

func newProductService(client *mongo.Client) *productService {
	return &productService{
		collection: client.Database("productsdb").Collection("products"),
	}
}

func (ps productService) ProductsGet(ctx context.Context) (openapi.ImplResponse, error) {
	cursor, err := ps.collection.Find(ctx, nil)
	if err != nil {
		panic(err)
	}

	var result []openapi.Product
	cursor.All(ctx, &result)

	return openapi.Response(http.StatusOK, result), nil
}

func (ps productService) ProductsIdDelete(_ context.Context, _ int32) (openapi.ImplResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (ps productService) ProductsIdGet(_ context.Context, _ int32) (openapi.ImplResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (ps productService) ProductsPost(ctx context.Context, p openapi.Product) (openapi.ImplResponse, error) {
	_, err := ps.collection.InsertOne(ctx, p)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}
