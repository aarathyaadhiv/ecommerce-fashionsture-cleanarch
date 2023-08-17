package main

import (
	"log"

	docs "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/cmd/api/docs"
	config "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/config"
	di "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/di"
)

func main() {
	docs.SwaggerInfo.Title = "MoviesGo - E-commerce"
	docs.SwaggerInfo.Description = "MoviesGo - E-commerce"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
