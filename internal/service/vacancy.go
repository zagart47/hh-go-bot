package service

import (
	"context"
	"encoding/json"
	"fmt"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/entity"
	"log"
	"strings"
)

type VacancyService struct {
	vacancier Vacancier
	converter Converter
	requester Requester
	messenger Messenger
}

func NewVacancyService(converter Converter, requestService RequestService, messenger Messenger) VacancyService {
	return VacancyService{
		vacancier: VacancyService{},
		converter: converter,
		requester: requestService,
		messenger: messenger,
	}
}

const GotResponse = "got_response"
const GotInvitation = "got_invitation"
const GotRejection = "got_rejection"

// All ищет все вакансии согласно строке поиска
func (s VacancyService) All(ctx context.Context, ch chan []string) {
	listMap := make(map[string]entity.Vacancy)
	vacancies := entity.NewVacancies()
	var link string
	for i := 0; ; i++ {
		link = fmt.Sprintf("https://api.hh.ru/vacancies?text=golang&id=publication_time&page=%d&per_page=100", i)
		body := s.requester.Request(ctx, link)
		err := json.Unmarshal(body, &vacancies)
		if err != nil {
			fmt.Println("Ошибка при десериализации ответа:", err)
		}
		for _, vacancy := range vacancies.Items {
			if strings.Contains(strings.ToLower(vacancy.Name), "go") {
				vacancy.Applied = s.vacancier.CheckRelations(ctx, vacancy.Relations)
				listMap[fmt.Sprintf("%s%s", vacancy.PublishedAt, vacancy.Id)] = vacancy
			}
		}
		if vacancies.Pages == i {
			break
		}
	}
	raw := s.converter.Convert(ctx, listMap)
	ch <- s.messenger.Message(ctx, raw)
}

// Similar нужен для поиска подходящих к резюме вакансий
func (s VacancyService) Similar(ctx context.Context, ch chan []string) {
	listMap := make(map[string]entity.Vacancy)
	vacancies := entity.NewVacancies()
	cfg, err := config.All()
	if err != nil {
		log.Print(err)
	}
	for i := 0; ; i++ {
		link := fmt.Sprintf("https://api.hh.ru/resumes/%s/similar_vacancies?id=publication_time&page=%d&per_page=100", cfg.Api.ResumeID, i)
		body := s.requester.Request(ctx, link)
		err = json.Unmarshal(body, &vacancies)
		if err != nil {
			fmt.Println("Ошибка при десериализации ответа:", err)
		}
		for _, vacancy := range vacancies.Items {
			if strings.Contains(strings.ToLower(vacancy.Name), "go") {
				vacancy.Applied = s.vacancier.CheckRelations(ctx, vacancy.Relations)
				listMap[fmt.Sprintf("%s%s", vacancy.PublishedAt, vacancy.Id)] = vacancy

			}
		}
		if vacancies.Pages == i {
			break
		}
	}
	raw := s.converter.Convert(ctx, listMap)
	ch <- s.messenger.Message(ctx, raw)
}

// CheckRelations нужен для проверки откликов на вакансии и установки эмодзи рядом с заголовком вакансии
func (s VacancyService) CheckRelations(ctx context.Context, ss []string) (r rune) {
	for _, v := range ss {
		switch v {
		case GotResponse:
			r = '\u26a0'
		case GotInvitation:
			r = '\u2705'
		case GotRejection:
			r = '\u274c'
		default:
			r = 0
		}
	}
	return
}
