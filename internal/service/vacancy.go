package service

import (
	"context"
	"encoding/json"
	"fmt"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/consts"
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

func (vs VacancyService) Vacancy(ctx context.Context, s string, chV chan []string) {
	listMap := make(map[string]entity.Vacancy)
	vacancies := entity.NewVacancies()
	var link string
	cfg, err := config.All()
	if err != nil {
		log.Println(err)
	}
	for i := 0; ; i++ {
		if strings.Contains(s, "similar_vacancies") {
			link = fmt.Sprintf(s, cfg.Api.ResumeID, i)
		} else {
			link = fmt.Sprintf(s, i)
		}
		ch := make(chan []byte)
		go vs.requester.Do(ctx, link, ch)
		body := <-ch
		err := json.Unmarshal(body, &vacancies)
		if err != nil {
			fmt.Println("Ошибка при десериализации ответа:", err)
		}
		for _, vacancy := range vacancies.Items {
			if strings.Contains(strings.ToLower(vacancy.Name), "go") {
				vacancy.Icon = vs.vacancier.CheckRelations(vacancy.Relations)
				listMap[fmt.Sprintf("%s%s", vacancy.PublishedAt, vacancy.Id)] = vacancy
			}
		}
		if vacancies.Pages == i {
			break
		}
	}
	vacanciesSlice := vs.converter.Convert(listMap)
	chV <- vs.messenger.MakeMessage(vacanciesSlice)
}

func (vs VacancyService) CheckRelations(ss []string) (r rune) {
	for _, v := range ss {
		switch v {
		case consts.GotResponse:
			r = '\u26a0'
		case consts.GotInvitation:
			r = '\u2705'
		case consts.GotRejection:
			r = '\u274c'
		default:
			r = 0
		}
	}
	return
}
