package main

//go:generate go run github.com/99designs/gqlgen generate

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/digiz3d/graphgogen/graph"
	"github.com/digiz3d/graphgogen/graph/generated"
	"github.com/digiz3d/graphgogen/graph/model"
	goredis "github.com/go-redis/redis/v9"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	redis := goredis.NewClient(&goredis.Options{Addr: redisAddr, ReadTimeout: time.Second * 60})

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ShowsRepository: make(map[string]*model.Show),
		UsersRepository: make(map[string]*model.User),
		Redis:           redis,
	}}))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})

	http.Handle("/playground", playground.Handler("Playground", "/graphql"))
	http.Handle("/graphql", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
