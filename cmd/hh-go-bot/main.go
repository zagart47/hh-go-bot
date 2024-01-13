package main

import (
	"context"
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/service"
	"hh-go-bot/transport/http"
	"hh-go-bot/transport/http/handler"
	"hh-go-bot/transport/telegrambot"
	"log"
)

func main() {
	services := service.NewServices()
	var f *string
	f = flag.String("d", "bot", "delivery using")
	flag.Parse()

	switch *f {
	case consts.BOT:
		config.All.SetMode(consts.BOT)
		bot, err := tgbotapi.NewBotAPI(config.All.Bot.Token)
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
	case consts.HTTP:
		config.All.SetMode(consts.HTTP)
		handlers := handler.NewHandler(&services)
		srv := http.NewServer(config.All.HTTP.Host, handlers.Init())
		err := srv.Run()
		if err != nil {
			log.Fatal(err)
		}
		ctx, shutdown := context.WithTimeout(context.Background(), consts.Timeout)
		defer shutdown()
		if err := srv.Stop(ctx); err != nil {
			log.Println("failed to stop server")
		}
	}
}
