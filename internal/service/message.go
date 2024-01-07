package service

import (
	"context"
	"fmt"
	"hh-go-bot/internal/entity"
)

type MessageService struct {
	messenger Messenger
}

func NewMessageService() MessageService {
	return MessageService{}
}

// Message делит список вакансий на массив по 40 вакансий
func (s MessageService) Message(ctx context.Context, vacancies entity.Vacancies) []string {
	var message string
	var messages []string
	var vacancyCount int
	for _, v := range vacancies.Items {
		message = fmt.Sprintf("%s\n%c %s | %s - %s", message, v.Applied, v.Employer.Name, v.Name, v.AlternateUrl)
		vacancyCount++
		if vacancyCount == 40 {
			messages = append(messages, message)
			vacancyCount = 0
			message = ""
		}
	}
	if message != "" {
		messages = append(messages, message)
		return messages
	}
	return nil
}
