package main

import (
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/service"
	"hh-go-bot/transport/telegrambot"
	"log"
)

func main() {
	cfg, err := config.All()
	if err != nil {
		log.Fatal(err)
	}
	services := service.NewServices()
	bot, err := telegrambot.NewBot(cfg.Bot.Token, services)
	if err != nil {
		log.Fatal(err)
	}
	err = bot.Start()
	if err != nil {
		return
	}
}
