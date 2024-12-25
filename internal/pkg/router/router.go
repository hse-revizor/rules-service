package router

import (
	docs "github.com/hse-revizor/rules-service/docs"
	"github.com/gin-gonic/gin"
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
	docs.SwaggerInfo.BasePath = "/api/v1"
	router := api.Group("/api/v1")
	{
		rules := router.Group("/rules")
		{
			rules.GET("/", h.GetRule)
			rules.GET("/all", h.GetRules)
			rules.POST("/", h.CreateRule)
			rules.DELETE("/", h.DeleteRule)
			rules.PUT("/", h.UpdateRule)
		}
	}
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
