package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"hh-go-bot/internal/entity"
	"log"
	"net/http"
	"time"
)

func (h Handler) initVacancyRoutes(api *gin.RouterGroup) {
	vacancy := api.Group("/vacancy")
	vacancy.GET("/all", h.Vacancy)
}

func (h Handler) Vacancy(c *gin.Context) {
	ch := make(chan any)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	go h.services.Vacancier.All(ctx, ch)

	select {
	case <-ctx.Done():
		log.Fatal("timeout")
	case j := <-ch:
		text := j.(entity.Vacancies)
		c.JSON(http.StatusOK, text)
	}
}
