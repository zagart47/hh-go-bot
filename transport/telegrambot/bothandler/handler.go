package bothandler

import (
	"hh-go-bot/transport/telegrambot"
)

type Bot struct {
	bot *telegrambot.BotService
}

func NewHandler(bot telegrambot.BotService) Bot {
	return Bot{
		bot: &bot,
	}
}

func (b *Bot) Init() {
	b.bot.Bot.Handle("/jobs", b.AllVacancies)
	b.bot.Bot.Handle("/similar", b.SimilarVacancies)
	b.bot.Bot.Handle("/resume", b.Resume)
}
