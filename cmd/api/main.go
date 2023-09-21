package main

import (
	"log"
	

	docs "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/cmd/api/docs"
	config "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/config"
	di "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/di"
	"github.com/joho/godotenv"
	
)

func main() {
	docs.SwaggerInfo.Title = "FashionStore - E-commerce"
	docs.SwaggerInfo.Description = "FashionStore - E-commerce"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	err:=godotenv.Load()

	if err!=nil{
		log.Fatal("cannot load env:",err)
	}
	
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
