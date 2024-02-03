package telegrambot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/service"
	"hh-go-bot/pkg/logger"
)

type BotAPI interface {
	Send(tgbotapi.Chattable) (tgbotapi.Message, error)
	GetUpdatesChan(tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
}

type BotService struct {
	bot      BotAPI
	services service.Services
}

func NewBotService(bot *tgbotapi.BotAPI, services service.Services) BotService {
	return BotService{
		bot:      bot,
		services: services,
	}
}

func (b BotService) Send(msg tgbotapi.Chattable) (tgbotapi.Message, error) {
	_, err := b.bot.Send(msg)
	if err != nil {
		return tgbotapi.Message{}, err
	}
	return tgbotapi.Message{}, nil
}

func (b BotService) GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel {
	return b.bot.GetUpdatesChan(config)
}

func (b BotService) Echo() error {
	config.All.SetMode(consts.BOT)
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 15
	updates := b.GetUpdatesChan(update)
	var text []string
	for u := range updates {
		ctx, cancel := context.WithTimeout(context.Background(), consts.Timeout)
		defer cancel()
		if u.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(u.Message.Chat.ID, u.Message.Text)
		msg.DisableWebPagePreview = true

		switch u.Message.Command() {

		case "similar":
			v, err := b.services.Vacancy.Similar(ctx)
			if err != nil {
				logger.Log.Error("similar vacancies getting error", err.Error())
			}
			text = VacancyMessage(v)

		case "jobs":
			v, err := b.services.Vacancy.All(ctx)
			if err != nil {
				logger.Log.Error("all vacancies getting error", err.Error())
			}
			text = VacancyMessage(v)
		case "resume":
			r, err := b.services.Resume.Get(ctx)
			if err != nil {
				logger.Log.Error("resume getting error", err.Error())
			}
			text = ResumeMessage(r)

		default:
			text = []string{"I don't know that command"}
		}
		for _, v := range text {
			msg.Text = v
			_, err := b.Send(msg)
			if err != nil {
				logger.Log.Warn("cannot send message by bot", "error", err.Error())
			}
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

func ResumeMessage(r entity.Resume) []string {
	var text []string
	for _, v := range r.Items {
		text = append(text, v.ID)
	}
	return text
}
