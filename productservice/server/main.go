package main

import (
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/worming004/fullsystemtest/productservice"
)

type server struct{}

func newServer() (server, error) {
	return server{}, nil
}

var allProducts = productservice.ProductSlice{
	productservice.Product{Id: 0, Name: "Chaussures", Price: 20},
	productservice.Product{Id: 1, Name: "Gants", Price: 10},
	productservice.Product{Id: 2, Name: "Chapeau", Price: 15},
}

func (s server) GetProducts(ctx echo.Context) error {
	t := struct {
		Products []productservice.Product
	}{
		Products: allProducts,
	}
	return ctx.JSON(http.StatusOK, t)

}
func (s server) GetProductsId(ctx echo.Context, id int) error {
	productFound := allProducts.Where(func(p productservice.Product) bool {
		return p.Id == id
	})

	if len(productFound) == 0 {
		return ctx.JSON(http.StatusNotFound, "Not found")
	}
	return ctx.JSON(http.StatusOK, productFound.First())
}

func main() {
	swagger, err := productservice.GetSwagger()
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil
	fn := middleware.OapiRequestValidator(swagger)

	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(fn)

	srv, err := newServer()
	if err != nil {
		panic(err)
	}
	productservice.RegisterHandlers(e, srv)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
