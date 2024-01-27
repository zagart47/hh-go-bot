package usecase

import (
	"context"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/repository"
)

type Vacancy interface {
	Get(context.Context, string) (map[string]entity.Vacancy, error)
	respondStatus([]string) rune
}

type Requester interface {
	Request(context.Context, string) []byte
}

type Resumes interface {
	GetResume(context.Context) (entity.Resume, error)
}

type Sorter interface {
	Sort(context.Context, map[string]entity.Vacancy) entity.Vacancies
}

type User interface {
	Name() string
	ID() int64
}

type Usecases struct {
	Vacancy       Vacancy
	Request       Requester
	Resumes       Resumes
	VacanciesRepo repository.Repositories
	User          User
	Sorter        Sorter
}

func NewUsecases() Usecases {
	requestUsecase := NewRequestUsecase()
	vacancyUsecase := NewVacancyUsecase(requestUsecase)
	resumeUsecase := NewResumeUsecase(requestUsecase)
	userUsecase := NewUserUsecase()
	sortUsecase := NewSortUsecase()
	return Usecases{
		Vacancy:       vacancyUsecase,
		Request:       requestUsecase,
		Resumes:       resumeUsecase,
		VacanciesRepo: repository.Repositories{},
		User:          userUsecase,
		Sorter:        sortUsecase,
	}
}
