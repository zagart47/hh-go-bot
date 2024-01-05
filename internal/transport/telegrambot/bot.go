package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/service"
	"time"
)

func New(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return bot, nil
}

func Start(bot *tgbotapi.BotAPI) (err error) {
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)

	jobService := service.NewService()
	vacancies := entity.NewVacancyList()

	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.DisableWebPagePreview = true
		switch update.Message.Command() {
		case "similar":
			vacancies = jobService.Similar()
		case "jobs":
			vacancies = jobService.All()
		default:
			msg.Text = "I don't know that command"
		}
		vacs := jobService.Message(vacancies)
		for _, v := range vacs {
			msg.Text = v
			if _, err = bot.Send(msg); err != nil {
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	return
}
