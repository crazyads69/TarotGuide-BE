package main

import (
	"golang_template/router"
	"golang_template/util"

	_ "golang_template/docs"

	"github.com/rs/zerolog/log"
)

// @title           Golang Template Swagger Documentation
// @version         0.0.1
// @description     This is the Swagger documentation for golang template.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Khiem Le
// @contact.url    https://khiemle.dev
// @contact.email  khiemledev@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not load config")
		return
	}
	// Watch config file changes
	util.WatchConfig(&cfg)

	util.ConfigLogger(cfg)

	server, err := router.NewAPIServer()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not create router")
		return
	}

	server.SetupRouter()
	server.SetupSwagger(cfg.SwaggerURL)

	err = server.Start(cfg.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start server")
		return
	}

	log.Info().Msgf("Listening and serving HTTP on %s", cfg.HTTPServerAddress)
}
