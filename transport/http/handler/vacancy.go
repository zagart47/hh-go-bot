package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"hh-go-bot/internal/config"
	"log"
	"net/http"
	"time"
)

func (h Handler) initVacancyRoutes(api *gin.RouterGroup) {
	vacancy := api.Group("/vacancy")
	vacancy.GET("/all", h.Vacancy)
}

func (h Handler) Vacancy(c *gin.Context) {
	ch := make(chan []string)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	go h.services.Vacancier.Vacancy(ctx, config.AllVacancies, ch)

	select {
	case <-ctx.Done():
		log.Fatal("timeout")
	case j := <-ch:
		c.JSON(http.StatusOK, j)
	}
}
