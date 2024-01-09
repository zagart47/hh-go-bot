package service

import (
	"context"
	"time"
)

type ContextService struct {
	context Context
}

func NewContextService() *ContextService {
	return &ContextService{}
}

func (c ContextService) WithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*20)
}
