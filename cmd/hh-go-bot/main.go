package main

import (
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/service"
	"hh-go-bot/transport/http"
	"hh-go-bot/transport/http/handler"
	"hh-go-bot/transport/telegrambot"
	"log"
)

const (
	BOT  = "bot"
	HTTP = "http"
)

func main() {
	cfg, err := config.All()
	if err != nil {
		log.Fatal(err)
	}
	services := service.NewServices()
	var f *string
	f = flag.String("d", "bot", "delivery using")
	flag.Parse()

	switch *f {
	case BOT:
		bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
		if err != nil {
			log.Fatal(err)
		}
		bot.Debug = true
		log.Printf("Authorized on account %s", bot.Self.UserName)
		bs := telegrambot.NewBotService(bot, services)
		err = bs.Echo()
		if err != nil {
			log.Fatal(err)
		}
	case HTTP:
		handlers := handler.NewHandler(&services)
		srv := http.NewServer(cfg.HTTP.Host, handlers.Init())
		err := srv.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
