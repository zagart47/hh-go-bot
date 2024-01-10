package service

import (
	"fmt"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
)

type MessageService struct {
	messenger Messenger
}

func NewMessageService() MessageService {
	return MessageService{}
}

// MakeMessage делит список вакансий на массив по 40 вакансий в каждом элементе,
// чтобы уложиться в лимит символов (4096) в сообщении
func (s MessageService) MakeMessage(vacancies entity.Vacancies) []string {
	var previousExp, message string
	var messages []string
	var vacancyCount int
	for _, v := range vacancies.Items {
		if previousExp == "" {
			message = fmt.Sprintf("%s\n%s: %s\n", message, consts.RequireExp, v.Experience.Name)
			previousExp = consts.NoExperience
		}
		if previousExp != v.Experience.ID {
			messages = append(messages, message)
			vacancyCount = 0
			message = fmt.Sprintf("\n%s: %s\n", consts.RequireExp, v.Experience.Name)
			previousExp = v.Experience.ID
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
			message = fmt.Sprintf("%s: %s (продолжение)\n", consts.RequireExp, v.Experience.Name)
		}
	}
	if message != "" {
		messages = append(messages, message)
		return messages
	} else {
		return messages
	}
}
