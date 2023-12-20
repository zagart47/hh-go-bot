package main

import (
	"hh-go-bot/bot"
	"hh-go-bot/config"
)

func main() {
	cfg := config.New()
	b := bot.NewBot(cfg)
	b.Start()
}
