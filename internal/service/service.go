package service

import (
	"context"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/repository"
	"hh-go-bot/internal/usecase"
)

type Vacancy interface {
	All(context.Context) (entity.Vacancies, error)
	Similar(context.Context) (entity.Vacancies, error)
}

type Resume interface {
	Get(context.Context) (entity.Resume, error)
}

type Services struct {
	Vacancy Vacancy
	Resume  Resume
}

func NewServices(repositories repository.Repositories) Services {
	useCases := usecase.NewUsecases(repositories)
	vacancyService := NewVacancyService(useCases)
	resumeService := NewResumeService(useCases)
	return Services{
		Vacancy: vacancyService,
		Resume:  resumeService,
	}
}
