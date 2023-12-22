package main

import (
	"hh-go-bot/internal/bot"
	"hh-go-bot/internal/config"
)

func main() {
	cfg := config.New()
	b := bot.NewBot(cfg)
	b.Start()
}
