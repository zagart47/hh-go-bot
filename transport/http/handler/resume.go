package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"hh-go-bot/internal/consts"
	"log"
	"net/http"
	"time"
)

func (h Handler) initResumeRoutes(api *gin.RouterGroup) {
	vacancy := api.Group("/resume")
	vacancy.GET("/mine", h.Resume)
}

func (h Handler) Resume(c *gin.Context) {
	ch := make(chan []string)
	ctx, cancel := context.WithTimeout(context.Background(), consts.Timeout*time.Second)
	defer cancel()
	go h.services.Resumes.MyResume(ctx, ch)

	select {
	case <-ctx.Done():
		log.Fatal("timeout")
	case resumes := <-ch:
		c.JSON(http.StatusOK, resumes)
	}
}
