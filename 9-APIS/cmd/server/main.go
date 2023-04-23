package main

import (
	"net/http"

	"github.com/devfullcycle/goexpert/9-APIS/configs"
	"github.com/devfullcycle/goexpert/9-APIS/internal/entity"
	"github.com/devfullcycle/goexpert/9-APIS/internal/infra/database"
	"github.com/devfullcycle/goexpert/9-APIS/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	productHandler := handlers.NewProductHandler(database.NewProduct(db))
	userHandler := handlers.NewUserHandler(database.NewUser(db));

	// Router Chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHandler.CreateUser)

	http.ListenAndServe(":8000", r)
}
