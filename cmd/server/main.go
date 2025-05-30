package main

import (
	"AiCheto/docs"
	"AiCheto/internal/delivery/api"
	"AiCheto/internal/repositories"
	"AiCheto/internal/usecases"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"log"
	"net/http"
)

// @title           Quotes API
// @version         1.0
// @description     Simple CRUD service for inspirational quotes
// @BasePath        /
func main() {
	repo := repositories.New()
	uc := usecases.NewQuoteUsecase(repo)
	h := api.NewQuoteHandler(uc)

	r := mux.NewRouter()

	h.Register(r)

	docs.SwaggerInfo.Title = "Quotes API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "Simple CRUD service for inspirational quotes"
	docs.SwaggerInfo.BasePath = "/"

	// Swagger - http://localhost:8080/swagger/index.html
	r.PathPrefix("/swagger/").
		Handler(httpSwagger.WrapHandler)

	log.Println("Quotes service started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
