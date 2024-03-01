package bothandler

import (
	"gopkg.in/telebot.v3"
	"hh-go-bot/transport/telegrambot"
)

type Bot struct {
	*telegrambot.BotService
}

func NewHandler(bot telegrambot.BotService) Bot {
	return Bot{
		&bot,
	}
}

func (b *Bot) Init() {
	b.Bot.Handle("/jobs", b.AllVacancies)
	b.Bot.Handle("/similar", b.SimilarVacancies)
	b.Bot.Handle("/resume", b.Resume)
	b.Bot.Handle(telebot.OnText, b.ParseCmd)
}
