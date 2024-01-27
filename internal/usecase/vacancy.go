package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/repository"
	"hh-go-bot/pkg/logger"
)

const (
	gotResponse   = "got_response"
	gotInvitation = "got_invitation"
	gotRejection  = "got_rejection"

	gotResponseIcon   = 9888  // '\u26a0'
	gotInvitationIcon = 9989  // '\u2705'
	gotRejectionIcon  = 10060 // '\u274c'
)

type VacancyUsecase struct {
	vacancy       Vacancy
	request       RequestUsecase
	vacanciesRepo repository.Repositories
}

func NewVacancyUsecase(req RequestUsecase) VacancyUsecase {
	return VacancyUsecase{
		vacancy: VacancyUsecase{},
		request: req,
	}
}

func (vs VacancyUsecase) Get(ctx context.Context, s string) (map[string]entity.Vacancy, error) {
	var m map[string]entity.Vacancy
	vacancies := entity.NewVacancies()
	var link string

	for i := 0; ; i++ {
		if s == consts.SimilarVacanciesLink {
			link = fmt.Sprintf(s, config.All.Api.ResumeID, i)
		} else {
			link = fmt.Sprintf(s, i)
		}

		body := vs.request.Request(ctx, link)
		err := json.Unmarshal(body, &vacancies)
		if err != nil {
			logger.Log.Warn("unmarshalling error", err.Error())
		}

		if m == nil {
			m = make(map[string]entity.Vacancy, vacancies.Found)
		}

		for _, v := range vacancies.Items {
			v.Icon = vs.vacancy.respondStatus(v.Relations)
			m[v.Id] = v
		}

		if vacancies.Pages == vacancies.Page {
			break
		}
	}
	return m, nil
}

// respondStatus для вставки иконки статуса отклика на вакансию
func (vs VacancyUsecase) respondStatus(s []string) (r rune) {
	for _, v := range s {
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
