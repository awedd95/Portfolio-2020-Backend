package main

import (
	"log"
	"net/http"
	"os"
	"server/graph"
	"server/graph/generated"
    "server/auth"
    database	"server/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
    "github.com/go-pg/pg/v10"
    "github.com/go-chi/chi"
  	"github.com/go-chi/chi/middleware"
)

const defaultPort = ":8888"

var db * pg.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
    dbpass := os.Getenv("PASSWORD")
    db := pg.Connect(&pg.Options{
        Addr:     ":5432",
        User: "postgres",
        Password: dbpass,
        Database: "backend",
    })
    defer db.Close()
    err := database.CreateSchema(db)
	if err != nil {
        log.Printf(dbpass)
		panic(err)
	}

        schema := generated.NewExecutableSchema(generated.Config{
            Resolvers: &graph.Resolver{
                DB: db,
            },
        })

    r := chi.NewRouter()
  	r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(auth.UserCtx)
	r.Route("/graphql", func(r chi.Router) {
		srv := handler.NewDefaultServer(schema)

		r.Handle("/", srv)
	})

	gqlPlayground := playground.Handler("api-gateway", "/graphql")
	r.Get("/graphiql", gqlPlayground)
	http.ListenAndServe(port, r)

	log.Printf("connect to http://localhost%s/graphiql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
