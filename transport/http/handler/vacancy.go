package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"hh-go-bot/internal/consts"
	"net/http"
)

func (h Handler) initVacanciesRoutes(api *gin.RouterGroup) {
	vacancy := api.Group("/vacancy")
	vacancy.GET("/all", h.AllVacancies)
	vacancy.GET("/similar", h.SimilarVacancies)
}

func (h Handler) AllVacancies(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), consts.Timeout)
	defer cancel()

	vacancies, err := h.services.Vacancy.All(ctx)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"all vacancies calling error": err.Error()})
	}
	c.JSON(http.StatusOK, vacancies)
}

func (h Handler) SimilarVacancies(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), consts.Timeout)
	defer cancel()

	vacancies, err := h.services.Vacancy.Similar(ctx)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"similar vacancies calling error": err.Error()})
		c.Abort()
	}
	c.JSON(http.StatusOK, vacancies)
}
