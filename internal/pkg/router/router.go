package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	docs "github.com/hse-revizor/rules-service/docs"
	"github.com/hse-revizor/rules-service/internal/pkg/service/rule"
	"github.com/hse-revizor/rules-service/internal/utils/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	cfg     *config.Config
	service *rule.Service
}

func NewRouter(cfg *config.Config, service *rule.Service) *Handler {
	return &Handler{
		cfg:     cfg,
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	api := gin.New()

	api.Use(gin.Recovery())
	api.Use(gin.Logger())
	api.Use(cors.Default())

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	docs.SwaggerInfo.Title = "Projects Service API"
	docs.SwaggerInfo.Description = "API Documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8788"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router := api.Group("/api")
	{
		rules := router.Group("/rule")
		{
			rules.POST("", h.CreateRule)
			rules.GET("/:id", h.GetRule)
			rules.DELETE("/:id", h.DeleteRule)
		}

		policies := router.Group("/policy")
		{
			policies.POST("", h.CreatePolicy)
			policies.GET("/:id", h.GetPolicy)
			policies.DELETE("/:id", h.DeletePolicy)
		}
	}

	return api
}
