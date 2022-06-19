package main

//go:generate go run github.com/99designs/gqlgen generate

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/digiz3d/graphgogen/graph"
	"github.com/digiz3d/graphgogen/graph/generated"
	"github.com/digiz3d/graphgogen/graph/model"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ShowsRepository: make(map[string]*model.Show),
		UsersRepository: make(map[string]*model.User),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
