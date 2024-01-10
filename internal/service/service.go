package service

import (
	"context"
	"hh-go-bot/internal/entity"
)

type Vacancier interface {
	Vacancy(context.Context, string, chan []string)
	CheckRelations([]string) rune
}

type Requester interface {
	Do(context.Context, string, chan []byte)
}

type Resumes interface {
	MyResume(context.Context, chan []string)
}

type Messenger interface {
	MakeMessage(entity.Vacancies) []string
}

type Converter interface {
	Convert(map[string]entity.Vacancy) entity.Vacancies
}
type Services struct {
	Vacancier Vacancier
	Requester Requester
	Resumes   Resumes
	Converter Converter
	Messenger Messenger
}

func NewServices() Services {
	converterService := NewConverterService()
	requestService := NewRequestService()
	messageService := NewMessageService()
	vacancyService := NewVacancyService(converterService, requestService, messageService)
	resumeService := NewResumeService(requestService)
	return Services{
		Vacancier: vacancyService,
		Requester: requestService,
		Resumes:   resumeService,
		Converter: converterService,
		Messenger: messageService,
	}
}
