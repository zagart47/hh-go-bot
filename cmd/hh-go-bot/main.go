package main

import (
	"context"
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/repository"
	"hh-go-bot/internal/repository/cache"
	"hh-go-bot/internal/repository/postgresql"
	"hh-go-bot/internal/service"
	"hh-go-bot/transport/http"
	"hh-go-bot/transport/http/handler"
	"hh-go-bot/transport/telegrambot"
	"log"
)

func main() {
	redisClient := cache.NewRedisClient(config.All.Redis.Host, config.All.Redis.Pwd)
	redisService := cache.NewRedisService(redisClient)

	db, err := postgresql.New()
	repos := repository.NewRepositories(db, redisService)
	services := service.NewServices(repos)
	if err != nil {
		log.Printf("db pool create error %s", err)
	}

	var f *string
	f = flag.String("d", "http", "delivery using")
	flag.Parse()

	switch *f {
	case consts.BOT:
		bot, err := tgbotapi.NewBotAPI(config.All.Bot.Token)
		if err != nil {
			log.Fatal(err)
		}
		bot.Debug = true
		_ = entity.NewUser(bot.Self.UserName, bot.Self.ID)

		log.Printf("Authorized on account %s", bot.Self.UserName)
		bs := telegrambot.NewBotService(bot, services)
		err = bs.Echo()
		if err != nil {
			log.Fatal(err)
		}
	case consts.HTTP:
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
