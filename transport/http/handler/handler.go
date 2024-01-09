package handler

import (
	"github.com/gin-gonic/gin"
	"hh-go-bot/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h Handler) Init() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "hello!")
	})
	h.initApi(router)
	return router
}

func (h Handler) initApi(router *gin.Engine) {
	api := router.Group("")
	h.initVacancyRoutes(api)
}
