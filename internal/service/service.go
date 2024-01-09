package service

import (
	"context"
	"hh-go-bot/internal/entity"
)

type Vacancier interface {
	All(context.Context, chan any)
	Similar(context.Context, chan any)
	CheckRelations([]string) rune
}

type Requester interface {
	Do(context.Context, string) []byte
}

type Resumes interface {
	MyResume(context.Context, chan any)
}

type Converter interface {
	Convert(map[string]entity.Vacancy) entity.Vacancies
}

type Messenger interface {
	Message(entity.Vacancies) []string
}

type Responder interface {
	Respond(context.Context, string)
}

type Services struct {
	Vacancier Vacancier
	Requester Requester
	Resumes   Resumes
	Converter Converter
	Messenger Messenger
	Responder Responder
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
