package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
	"log"
	"net/http"
)

func (h Handler) initAllVacanciesRoutes(api *gin.RouterGroup) {
	vacancy := api.Group("/vacancy")
	vacancy.GET("/all", h.AllVacancies)
}

func (h Handler) AllVacancies(c *gin.Context) {
	ch := make(chan any)

	ctx, cancel := context.WithTimeout(context.Background(), consts.Timeout)
	defer cancel()
	go h.services.Vacancier.Vacancy(ctx, consts.AllVacanciesLink, ch)

	select {
	case <-ctx.Done():
		log.Fatal("timeout")
	case vacancies := <-ch:
		c.JSON(http.StatusOK, vacancies.(entity.Vacancies))
	}
}
