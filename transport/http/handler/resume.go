package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"hh-go-bot/internal/consts"
	"net/http"
)

func (h Handler) initResumeRoutes(api *gin.RouterGroup) {
	vacancy := api.Group("/resume")
	vacancy.GET("/mine", h.Resume)
}

func (h Handler) Resume(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), consts.Timeout)
	defer cancel()

	resumes, err := h.services.Resume.Get(ctx)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"resume getting error": err.Error()})
	}
	c.JSON(http.StatusOK, resumes)
}
