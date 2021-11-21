package main

import (
	"log"
	"net/http"
	"os"
	"test_task/graph"
	"test_task/graph/generated"
	"test_task/internal/repository"
	"test_task/internal/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatalf("No found .env")
	}

	db := pg.Connect(&pg.Options{
		Addr:     ":" + os.Getenv("port"),
		User:     os.Getenv("user"),
		Password: os.Getenv("password"),
		Database: os.Getenv("dbname"),
	})

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(service)}))

	router := chi.NewRouter()
	router.Use(graph.Middleware(db))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", ":8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
