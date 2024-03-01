package bothandler

import (
	"context"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	selector   = &tele.ReplyMarkup{}
	btnRespond = selector.Data("Откликнуться", "prev")
	btnClose   = selector.Data("Закрыть", "close")
)

func (b *Bot) AllVacancies(c tele.Context) error {
	v, err := b.Services.Vacancy.All(context.Background())
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

func (b *Bot) ParseCmd(c tele.Context) error {
	switch {
	case strings.HasPrefix(c.Text(), "/o"):
		return b.OpenVacancy(c)
	default:
		return c.Send("Nothing")
	}

}

func (b *Bot) OpenVacancy(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := c.Text()[2:]
	one, err := b.Services.Vacancy.One(ctx, id)
	if err != nil {
		return err
	}
	text := OneVacancyMessage(one)
	selector.Inline(
		selector.Row(btnRespond),
		selector.Row(btnClose),
	)
	b.Bot.Handle(&btnRespond, func(c tele.Context) error {
		if err := b.MakeRespond(id); err != nil {
			return err
		}
		return c.Send("Отклик на вакансию направлен")
	})
	b.Bot.Handle(&btnClose, func(c tele.Context) error {
		return c.Delete()
	})
	return c.EditOrSend(text, &tele.SendOptions{DisableWebPagePreview: true}, selector)
}

func OneVacancyMessage(v entity.Vacancy) string {
	return fmt.Sprintf("%c\nКомпания: %s\nДолжность: %s\nСсылка: %s\nОпыт: %s\nКлючевые навыки: %s", v.Icon, v.Employer.Name, v.Name, v.AlternateUrl, v.Exp.Name, func() string {
		var list string
		for _, vac := range v.KeySkills {
			list = fmt.Sprintf("%s, %s", list, vac.Name)
		}
		_, i := utf8.DecodeRuneInString(list)
		return list[i:]
	}())
}

func (b *Bot) MakeRespond(id string) error {
	req, err := http.NewRequest(http.MethodPost, consts.RespondLink, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.All.Api.Bearer))
	formData := url.Values{}
	formData.Add("message", config.All.Bot.RespondText)
	formData.Add("resume_id", config.All.Api.ResumeID)
	formData.Add("vacancy_id", id)
	req.Body = io.NopCloser(strings.NewReader(formData.Encode()))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return err
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}
	return nil
}

func (b *Bot) SimilarVacancies(c tele.Context) error {
	v, err := b.Services.Vacancy.Similar(context.Background())
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
			message = fmt.Sprintf("%s\n%s | %s - /o%s", message, v.Employer.Name, v.Name, v.Id)
		} else {
			message = fmt.Sprintf("%s\n%c%s | %s - /o%s", message, v.Icon, v.Employer.Name, v.Name, v.Id)
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
