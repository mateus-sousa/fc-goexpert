package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/mateus-sousa/goexpert/7-api/configs"
	_ "github.com/mateus-sousa/goexpert/7-api/docs"
	"github.com/mateus-sousa/goexpert/7-api/internal/entity"
	"github.com/mateus-sousa/goexpert/7-api/internal/infra/database"
	"github.com/mateus-sousa/goexpert/7-api/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

// @title           Go Expert Example
// @version         1.0
// @description     Product API with authentication.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Mateus Silva
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Panda License
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg, err := configs.LoadConfig(".")
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

	userDb := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDb)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", cfg.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", cfg.JWTExpiresIn))
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})
	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)
}
