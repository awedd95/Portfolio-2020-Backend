package main

import (
	"log"
	"net/http"
	"os"
	"server/graph"
	"server/graph/generated"
    database	"server/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
    "github.com/go-pg/pg/v10"
)

const defaultPort = "8888"

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
		panic(err)
	}

        schema := generated.NewExecutableSchema(generated.Config{
            Resolvers: &graph.Resolver{
                DB: db,
            },
        })
	    srv := handler.NewDefaultServer(schema)

	http.Handle("/graphql", srv)

    gqlPlayground := playground.Handler("GraphQL playground", "/graphql")
	http.Handle("/graphiql", gqlPlayground)

	log.Printf("connect to http://localhost:%s/graphiql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
