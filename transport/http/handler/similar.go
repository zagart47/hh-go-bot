package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"hh-go-bot/internal/consts"
	"log"
	"net/http"
	"time"
)

func (h Handler) initSimilarVacancyRoutes(api *gin.RouterGroup) {
	vacancy := api.Group("/vacancy")
	vacancy.GET("/similar", h.Similar)
}

func (h Handler) Similar(c *gin.Context) {
	ch := make(chan []string)
	ctx, cancel := context.WithTimeout(context.Background(), consts.Timeout*time.Second)
	defer cancel()
	go h.services.Vacancier.Vacancy(ctx, consts.SimilarVacancies, ch)

	select {
	case <-ctx.Done():
		log.Fatal("timeout")
	case vacancies := <-ch:
		c.JSON(http.StatusOK, vacancies)
	}
}
