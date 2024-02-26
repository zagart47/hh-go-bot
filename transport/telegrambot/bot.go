package telegrambot

import (
	tele "gopkg.in/telebot.v3"
	"hh-go-bot/internal/service"
	"hh-go-bot/pkg/logger"
)

type BotService struct {
	Bot      tele.Bot
	Services service.Services
}

func NewBotService(pref tele.Settings, services service.Services) BotService {
	b, err := tele.NewBot(pref)
	if err != nil {
		logger.Log.Error("cannot create bot", err.Error())
	}
	return BotService{
		Bot:      *b,
		Services: services,
	}
}
