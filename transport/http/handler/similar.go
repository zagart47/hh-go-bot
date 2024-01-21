package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
	"log"
	"net/http"
)

func (h Handler) initSimilarVacanciesRoutes(api *gin.RouterGroup) {
	vacancy := api.Group("/vacancy")
	vacancy.GET("/similar", h.SimilarVacancies)
}

func (h Handler) SimilarVacancies(c *gin.Context) {
	ch := make(chan any)
	ctx, cancel := context.WithTimeout(context.Background(), consts.Timeout)
	defer cancel()
	go h.services.Vacancier.Vacancy(ctx, consts.SimilarVacanciesLink, ch)

	select {
	case <-ctx.Done():
		log.Fatal("timeout")
	case vacancies := <-ch:
		c.JSON(http.StatusOK, vacancies.(entity.Vacancies))
	}
}
