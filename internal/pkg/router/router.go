package router

import (
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
func (h *Handler) InitRoutes() {
	api := gin.New()
	docs.SwaggerInfo.BasePath = "/api"
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
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
