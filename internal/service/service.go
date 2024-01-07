package service

import (
	"context"
	"hh-go-bot/internal/entity"
)

type Vacancier interface {
	All(context.Context, chan []string)
	Similar(context.Context, chan []string)
	CheckRelations(context.Context, []string) rune
}

type Requester interface {
	Request(context.Context, string) []byte
}

type Resumes interface {
	MyResume(context.Context, chan []string)
}

type Converter interface {
	Convert(context.Context, map[string]entity.Vacancy) entity.Vacancies
}

type Messenger interface {
	Message(context.Context, entity.Vacancies) []string
}

type Responder interface {
	Respond(context.Context, string) entity.Vacancies
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
