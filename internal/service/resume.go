package service

import (
	"context"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/usecase"
)

type ResumeService struct {
	resume  Resume
	usecase usecase.Usecases
}

func NewResumeService(usecase usecase.Usecases) ResumeService {
	return ResumeService{
		resume:  ResumeService{},
		usecase: usecase,
	}
}

func (s ResumeService) Get(ctx context.Context) (entity.Resume, error) {
	return s.usecase.Resumes.GetResume(ctx)
}
