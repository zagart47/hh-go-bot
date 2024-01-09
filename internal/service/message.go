package service

import (
	"fmt"
	"hh-go-bot/internal/entity"
)

type MessageService struct {
	messenger Messenger
}

func NewMessageService() MessageService {
	return MessageService{}
}

// Message делит список вакансий на массив по 40 вакансий в каждом элементе,
// чтобы уложиться в лимит символов (4096) в сообщении
func (s MessageService) Message(vacancies entity.Vacancies) []string {
	var message string
	var messages []string
	var vacancyCount int
	for _, v := range vacancies.Items {
		message = fmt.Sprintf("%s\n%c %s | %s - %s", message, v.Icon, v.Employer.Name, v.Name, v.AlternateUrl)
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
	} else {
		return messages
	}
}
