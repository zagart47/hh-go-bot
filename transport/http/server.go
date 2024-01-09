package http

import (
	"context"
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
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
