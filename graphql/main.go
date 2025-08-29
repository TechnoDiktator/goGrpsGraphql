package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountURL string `envconfig:"ACCOUNT_SERVICE_URL" default:"localhost:50051"`
	CatalogURL string `envconfig:"CATALOG_SERVICE_URL" default:"localhost:50052"`
	OrderURL   string `envconfig:"ORDER_SERVICE_URL" default:"localhost:50053"`
}

func main() {

	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal("Failed to process env var: ", err)
	}

	s, err := NewGraphQLServer(cfg.AccountURL, cfg.CatalogURL, cfg.OrderURL)

	if err != nil {
		log.Fatal("Failed to create GraphQL server: ", err)
	}

	http.Handle("/graphql", handler.New(s.ToExecutableSchema()))

	//what is this playground
	//This is a GraphQL playground that allows you to interact with your GraphQL API through a web interface.
	http.Handle("/playground", playground.Handler("GraphQL playground by Tarang", "/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
