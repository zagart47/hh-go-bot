package main

import (
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/transport/telegrambot"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	b, err := telegrambot.New(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}
	if err = telegrambot.Start(b); err != nil {
		log.Fatal(err)
	}
}
