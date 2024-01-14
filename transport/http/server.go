package http

import (
	"context"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/consts"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(host string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    host,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	config.All.SetMode(consts.HTTP)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
