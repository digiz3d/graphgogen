package main

//go:generate go run github.com/99designs/gqlgen generate

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/digiz3d/graphgogen/graph"
	"github.com/digiz3d/graphgogen/graph/generated"
	"github.com/digiz3d/graphgogen/graph/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	gorilla "github.com/gorilla/websocket"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const defaultPort = "3000"

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cors := cors.New()
	app.Use(cors)

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ShowsRepository: make(map[string]*model.Show),
		UsersRepository: make(map[string]*model.User),
	}}))

	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: gorilla.Upgrader{CheckOrigin: func(r *http.Request) bool {
			return true
		}},
	})

	playgroundHttpHandler := playground.Handler("GraphQL playground", "/graphql")
	playgroundFastHttpHandler := fasthttpadaptor.NewFastHTTPHandlerFunc(playgroundHttpHandler)

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println(c.Method() + " " + c.Path())
		return c.Next()
	})

	app.Use("/playground", func(ctx *fiber.Ctx) error {
		playgroundFastHttpHandler(ctx.Context())
		return nil
	})

	adaptedSrv := fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	})
	app.Use("/graphql", func(ctx *fiber.Ctx) error {
		adaptedSrv(ctx.Context())
		return nil
	})

	app.Get("/ws", func(ctx *fiber.Ctx) error {
		adaptedSrv(ctx.Context())
		return nil
	})

	// app.Use("/ws", func(c *fiber.Ctx) error {
	// 	if websocket.IsWebSocketUpgrade(c) {
	// 		fmt.Println("ws upgrade")
	// 		c.Locals("allowed", true)
	// 		return c.Next()
	// 	}
	// 	fmt.Println("not upgraded")
	// 	return fiber.ErrUpgradeRequired
	// })

	// app.Get("/ws", websocket.New(func(c *websocket.Conn) {
	// 	// c.Locals is added to the *websocket.Conn
	// 	log.Println(c.Locals("allowed"))  // true
	// 	log.Println(c.Params("id"))       // 123
	// 	log.Println(c.Query("v"))         // 1.0
	// 	log.Println(c.Cookies("session")) // ""
	// 	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	// 	var (
	// 		mt  int
	// 		msg []byte
	// 		err error
	// 	)
	// 	for {
	// 		if mt, msg, err = c.ReadMessage(); err != nil {
	// 			log.Println("read:", err)
	// 			break
	// 		}
	// 		log.Printf("recv: %s", msg)
	// 		if err = c.WriteMessage(mt, msg); err != nil {
	// 			log.Println("write:", err)
	// 			break
	// 		}
	// 	}
	// }))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	app.Listen(":" + port)
}
