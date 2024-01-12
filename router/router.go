package router

import (
	"errors"
	controller "golang_template/controller"
	"golang_template/database"
	"golang_template/docs"
	"golang_template/helper"
	"golang_template/service"
	"golang_template/util"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db = database.DBConnect()

type APIServer struct {
	router *gin.Engine
}

func setupNotFoundPage(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		err := errors.New("not found")
		helper.NewHTTPError(
			ctx,
			http.StatusNotFound,
			err,
			err.Error(),
		)
	})
}

func setupChatRoutes(group *gin.RouterGroup) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Panic().Err(err).Msg("Failed to load config")
	}
	validator := validator.New()

	// Setup service
	chatRouter := controller.Controller{
		Config:        &config,
		ChatService:   service.ChatServiceImpl(),
		CardService:   service.CardServiceImpl(),
		PromptService: service.PromptServiceImpl(),
		Validator:     validator,
	}

	group.POST("/chat", chatRouter.Chat)

}

func NewAPIServer() (*APIServer, error) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Request-Id", "X-Requested-With"},
		MaxAge:           12 * time.Hour,
	}))

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
	setupChatRoutes(server.router.Group("/"))
}

func (server *APIServer) SetupSwagger(swaggerUrl string) {
	docs.SwaggerInfo.BasePath = "/"
	server.router.GET(swaggerUrl+"/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
