package main

import (
	"log"

	"jerseyhub/cmd/api/docs"
	config "jerseyhub/pkg/config"
	di "jerseyhub/pkg/di"
)

func main() {

	// // swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "JERSEYHUB"
	docs.SwaggerInfo.Description = "Here passion meets the fashion"
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
