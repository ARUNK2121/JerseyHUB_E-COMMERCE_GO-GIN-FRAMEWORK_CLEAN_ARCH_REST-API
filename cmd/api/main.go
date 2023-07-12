package main

import (
	"log"

	"jerseyhub/cmd/api/docs"
	config "jerseyhub/pkg/config"
	di "jerseyhub/pkg/di"

	"github.com/joho/godotenv"
)

// @SecurityDefinition BearerAuth
// @TokenUrl /auth/token

//	@securityDefinitions.Bearer		type apiKey
//	@securityDefinitions.Bearer		name Authorization
//	@securityDefinitions.Bearer		in header
//	@securityDefinitions.BasicAuth	type basic

func main() {

	// // swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "JERSEYHUB"
	docs.SwaggerInfo.Description = "Here passion meets the fashion,This is an online store for purchasing high quality jerseys of your favorite clubs.."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env file")
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
