package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/VituSuperMEg/graphql-fullcycle/graph"
	"github.com/VituSuperMEg/graphql-fullcycle/internal/database"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:3640@localhost/grpc-go?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDB,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
