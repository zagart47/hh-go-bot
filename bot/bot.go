package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hh-go-bot/config"
	"hh-go-bot/internal/model/job"
	job2 "hh-go-bot/internal/service/job"
	"log"
	"time"
)

type Bot struct {
	bot tgbotapi.BotAPI
}

func NewBot(cfg config.Config) Bot {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		panic(err.Error())
	}
	return Bot{bot: *bot}
}

func (tg Bot) Start() {
	tg.bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	updates := tg.bot.GetUpdatesChan(u)

	jobService := job2.NewMainService()
	vacancies := job.NewVacancyList()
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
			if _, err := tg.bot.Send(msg); err != nil {
				log.Panic(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}
