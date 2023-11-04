package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mateus-sousa/goexpert/7-api/configs"
	"github.com/mateus-sousa/goexpert/7-api/internal/entity"
	"github.com/mateus-sousa/goexpert/7-api/internal/infra/database"
	"github.com/mateus-sousa/goexpert/7-api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDb := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDb)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	http.ListenAndServe(":8000", r)
}
