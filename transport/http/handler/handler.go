package handler

import (
	"github.com/gin-gonic/gin"
	"hh-go-bot/internal/service"
)

type Handler struct {
	services service.Services
}

func NewHandler(services service.Services) Handler {
	return Handler{
		services: services,
	}
}

func (h Handler) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	api := router.Group("/")
	{
		h.initAllVacanciesRoutes(api)
		h.initSimilarVacanciesRoutes(api)
		h.initResumeRoutes(api)
	}
	return router
}
