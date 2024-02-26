package bothandler

import (
	"context"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
)

func (b *Bot) AllVacancies(c tele.Context) error {
	v, err := b.bot.Services.Vacancy.All(context.Background())
	if err != nil {
		return err
	}
	for _, vacs := range VacancyMessage(v) {
		err = c.Send(vacs, &tele.SendOptions{DisableWebPagePreview: true})
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) SimilarVacancies(c tele.Context) error {
	v, err := b.bot.Services.Vacancy.Similar(context.Background())
	if err != nil {
		return err
	}
	for _, vacs := range VacancyMessage(v) {
		err = c.Send(vacs, &tele.SendOptions{DisableWebPagePreview: true})
		if err != nil {
			return err
		}
	}
	return nil
}

const (
	ReqExp       = "Требуемый опыт"
	ContinuePage = "(продолжение)"
)

func VacancyMessage(vacancies entity.Vacancies) []string {
	var previousExp, message string
	var messages []string
	var vacancyCount int
	for _, v := range vacancies.Items {
		if previousExp == "" {
			message = fmt.Sprintf("%s\n%s: %s\n", message, ReqExp, v.Exp.Name)
			previousExp = consts.NoExp
		}
		if previousExp != v.Exp.ID {
			messages = append(messages, message)
			vacancyCount = 0
			message = fmt.Sprintf("\n%s: %s\n", ReqExp, v.Exp.Name)
			previousExp = v.Exp.ID
		}
		if v.Icon == 0 {
			message = fmt.Sprintf("%s\n%s | %s - %s", message, v.Employer.Name, v.Name, v.AlternateUrl)
		} else {
			message = fmt.Sprintf("%s\n%c%s | %s - %s", message, v.Icon, v.Employer.Name, v.Name, v.AlternateUrl)
		}
		vacancyCount++
		if vacancyCount == 40 {
			messages = append(messages, message)
			vacancyCount = 0
			message = fmt.Sprintf("%s: %s %s\n", ContinuePage, ReqExp, v.Exp.Name)
		}
	}
	if message != "" {
		messages = append(messages, message)
		return messages
	} else {
		return messages
	}
}
