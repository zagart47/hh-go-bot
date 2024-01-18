package telegrambot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/service"
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
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 15
	updates := b.GetUpdatesChan(u)
	for update := range updates {

		ctx, cancel := context.WithTimeout(context.Background(), consts.Timeout)
		defer cancel()
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.DisableWebPagePreview = true

		ch := make(chan any)
		switch update.Message.Command() {

		case "similar":
			go b.services.Vacancier.Vacancy(ctx, consts.SimilarVacancies, ch)

		case "jobs":
			go b.services.Vacancier.Vacancy(ctx, consts.AllVacancies, ch)

		case "resume":
			go b.services.Resumes.MyResume(ctx, ch)

		default:
			text := []string{"I don't know that command"}
			ch <- text
		}
		select {
		case <-ctx.Done():
			msg.Text = "20 sec timeout"
			_, err := b.Send(msg)
			if err != nil {
				return err
			}
		case raw := <-ch:
			text := raw.([]string)
			for _, v := range text {
				msg.Text = v
				_, err := b.Send(msg)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
