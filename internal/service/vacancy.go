package service

import (
	"context"
	"encoding/json"
	"fmt"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
	"strings"
)

const (
	gotResponse   = "got_response"
	gotInvitation = "got_invitation"
	gotRejection  = "got_rejection"

	gotResponseIcon   = 9888  // '\u26a0'
	gotInvitationIcon = 9989  // '\u2705'
	gotRejectionIcon  = 10060 // '\u274c'
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

func (vs VacancyService) Vacancy(ctx context.Context, s string, chV chan any) {
	listMap := make(map[string]entity.Vacancy)
	vacancies := entity.NewVacancies()
	var link string
	for i := 0; ; i++ {
		if strings.Contains(s, "similar_vacancies") {
			link = fmt.Sprintf(s, config.All.Api.ResumeID, i)
		} else {
			link = fmt.Sprintf(s, i)
		}
		ch := make(chan []byte)
		go vs.requester.doRequest(ctx, link, ch)
		body := <-ch
		err := json.Unmarshal(body, &vacancies)
		if err != nil {
			fmt.Println("Ошибка при десериализации ответа:", err)
		}
		for _, vacancy := range vacancies.Items {
			if strings.Contains(strings.ToLower(vacancy.Name), "go") {
				vacancy.Icon = vs.vacancier.checkRelations(vacancy.Relations)
				listMap[fmt.Sprintf("%s%s", vacancy.PublishedAt, vacancy.Id)] = vacancy
			}
		}
		if vacancies.Pages == i {
			break
		}
	}
	vacanciesSlice := vs.converter.convert(listMap)
	if config.All.Mode == consts.BOT {
		chV <- vs.messenger.makeMessage(vacanciesSlice)
	}
	if config.All.Mode == consts.HTTP {
		chV <- vacanciesSlice
	}
}

func (vs VacancyService) checkRelations(ss []string) (r rune) {
	for _, v := range ss {
		switch v {
		case gotResponse:
			r = gotResponseIcon
		case gotInvitation:
			r = gotInvitationIcon
		case gotRejection:
			r = gotRejectionIcon
		default:
			r = 0
		}
	}
	return
}
