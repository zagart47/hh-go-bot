package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/service"
)

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
	u.Timeout = 15
	updates := b.bot.GetUpdatesChan(u)
	for update := range updates {
		ctx, cancel := b.services.Context.WithTimeout()
		defer cancel()
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.DisableWebPagePreview = true

		ch := make(chan []string)
		switch update.Message.Command() {

		case "similar":
			go b.services.Vacancier.Vacancy(ctx, config.SimilarVacancies, ch)

		case "jobs":
			go b.services.Vacancier.Vacancy(ctx, config.AllVacancies, ch)

		case "resume":
			go b.services.Resumes.MyResume(ctx, ch)

		default:
			ch <- append([]string{}, "I don't know that command")
		}
		select {
		case <-ctx.Done():
			msg.Text = "20 sec timeout"
			_, err := b.bot.Send(msg)
			if err != nil {
				return err
			}
		case raw := <-ch:
			for _, v := range raw {
				msg.Text = v
				_, err := b.bot.Send(msg)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
