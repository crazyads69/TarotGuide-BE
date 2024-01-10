package router

import (
	"errors"
	"golang_template/docs"
	"golang_template/helper"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type APIServer struct {
	router *gin.Engine
}

func setupNotFoundPage(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		err := errors.New("Not Found")
		helper.NewHTTPError(
			ctx,
			http.StatusNotFound,
			err,
			err.Error(),
		)
	})
}

func NewAPIServer() (*APIServer, error) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   "false",
			"message": "Application is running",
		})
	})

	r.GET("/health_check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   "false",
			"message": "Server is running (Healthy)",
		})
	})

	router := &APIServer{
		router: r,
	}
	return router, nil
}

func (server *APIServer) Start(address string) error {
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return server.router.Run(address)
}

func (server *APIServer) SetupRouter() {
	setupNotFoundPage(server.router)
}

func (server *APIServer) SetupSwagger(swaggerUrl string) {
	docs.SwaggerInfo.BasePath = "/"
	server.router.GET(swaggerUrl+"/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
