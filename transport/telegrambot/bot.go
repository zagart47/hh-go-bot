package telegrambot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/service"
	"time"
)

/*type BotInterface interface {
	Start() error
	SendArrayMessage(*tgbotapi.BotAPI, tgbotapi.MessageConfig, []string)
}

type BotService struct {
	botInterface BotInterface
	botapi       *tgbotapi.BotAPI
	services     service.Services
}

func NewBot(token string, services service.Services) (*BotService, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &BotService{
		botapi:       bot,
		services:     services,
		botInterface: BotService{},
	}, nil
}

func (b BotService) Start() (err error) {
	b.botapi.Debug = true
	u := tgbotapi.NewUpdate(0)
	updates := b.botapi.GetUpdatesChan(u)
	for update := range updates {
		chVacancy := make(chan entity.Vacancies)
		chResume := make(chan []string)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.DisableWebPagePreview = true
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		switch update.Message.Command() {

		case "similar":
			go b.services.Vacancier.Similar(ctx, chVacancy)

		case "jobs":
			go b.services.Vacancier.All(ctx, chVacancy)

		case "resume":
			go b.services.Resumes.MyResume(ctx, chResume)

		default:
			go func() {
				chResume <- append([]string{}, "I don't know that command")
			}()
		}
		select {
		case <-ctx.Done():
			msg.Text = "20 sec timeout"
		case vacancies := <-chVacancy:
			raw := b.services.Messenger.Message(vacancies)
			b.botInterface.SendArrayMessage(b.botapi, msg, raw)
		case resume := <-chResume:
			b.botInterface.SendArrayMessage(b.botapi, msg, resume)
		}
		close(chResume)
		close(chVacancy)
	}
	return
}*/

func (b BotService) SendArrayMessage(bapi *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, m []string) {
	for _, v := range m {
		msg.Text = v
		if _, err := bapi.Send(msg); err != nil {
			fmt.Println("error")
		}
		//time.Sleep(500 * time.Millisecond)
	}
}

type BotAPI interface {
	Send(tgbotapi.Chattable) (tgbotapi.Message, error)
	GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
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

func (b BotService) Echo() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)
	for update := range updates {
		ch := make(chan any)
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.DisableWebPagePreview = true

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		switch update.Message.Command() {
		case "similar":
			go b.services.Vacancier.Similar(ctx, ch)

		case "jobs":
			go b.services.Vacancier.All(ctx, ch)

		case "resume":
			go b.services.Resumes.MyResume(ctx, ch)

		default:
			ch <- "I don't know that command"
		}
		select {
		case <-ctx.Done():
			msg.Text = "20 sec timeout"
		case raw := <-ch:
			switch raw.(type) {
			case []string:
				text := raw.([]string)
				for _, v := range text {
					msg.Text = v
					_, err := b.bot.Send(msg)
					if err != nil {
						return err
					}
				}
			case entity.Vacancies:
				text := raw.(entity.Vacancies)
				vacancies := b.services.Messenger.Message(text)
				for _, v := range vacancies {
					msg.Text = v
					_, err := b.bot.Send(msg)
					if err != nil {
						return err
					}
				}
			}
		}
		close(ch)
	}
	return nil
}
