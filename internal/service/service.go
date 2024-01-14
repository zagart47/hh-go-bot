package service

import (
	"context"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/repository/postgresql"
)

type Vacancier interface {
	Vacancy(context.Context, string, chan any)
	checkRelations([]string) rune
}

type Requester interface {
	doRequest(context.Context, string, chan []byte)
}

type Resumes interface {
	MyResume(context.Context, chan any)
}

type Messenger interface {
	makeMessage(entity.Vacancies) []string
}

type Converter interface {
	convert(map[string]entity.Vacancy) entity.Vacancies
}
type Services struct {
	Vacancier     Vacancier
	Requester     Requester
	Resumes       Resumes
	Converter     Converter
	Messenger     Messenger
	VacanciesRepo postgresql.Repositories
}

func NewServices(VacanciesRepo *postgresql.Repositories) Services {
	converterService := NewConverterService()
	requestService := NewRequestService()
	messageService := NewMessageService()
	vacancyService := NewVacancyService(converterService, requestService, messageService, VacanciesRepo)
	resumeService := NewResumeService(requestService)
	return Services{
		Vacancier:     vacancyService,
		Requester:     requestService,
		Resumes:       resumeService,
		Converter:     converterService,
		Messenger:     messageService,
		VacanciesRepo: *VacanciesRepo,
	}
}
