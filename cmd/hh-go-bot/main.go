package main

import (
	"context"
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/repository"
	"hh-go-bot/internal/repository/cache"
	"hh-go-bot/internal/repository/postgresql"
	"hh-go-bot/internal/service"
	"hh-go-bot/pkg/logger"
	"hh-go-bot/transport/http"
	"hh-go-bot/transport/http/handler"
	"hh-go-bot/transport/telegrambot"
	"log"
	"os"
)

func main() {
	redisClient := cache.NewRedisClient(config.All.Redis.Host, config.All.Redis.Pwd)
	redisService := cache.NewRedisService(redisClient)
	logger.Log.Info("client added", "redis", config.All.Redis.Host)

	db, err := postgresql.New()
	repos := repository.NewRepositories(db, redisService)
	if err != nil {
		logger.Log.Warn("repo creating error", "error", err)
		os.Exit(1)
	}

	services := service.NewServices(repos)

	f := flag.String("d", "http", "delivery using")
	flag.Parse()
	logger.Log.Info("starting application", "mode", f)

	switch *f {
	case consts.BOT:
		bot, err := tgbotapi.NewBotAPI(config.All.Bot.Token)
		if err != nil {
			logger.Log.Warn("new botapi creating error", "config", config.All.Bot.Token)
		}
		bot.Debug = true
		logger.Log.Info("bot debugging", "mode", bot.Debug)

		bs := telegrambot.NewBotService(bot, services)
		err = bs.Echo()
		if err != nil {
			logger.Log.Warn("bot starting error", "error", err)
			os.Exit(1)
		}
	case consts.HTTP:
		handlers := handler.NewHandler(services)
		logger.Log.Info("handlers initialized")
		srv := http.NewServer(config.All.HTTP.Host, handlers.Init())
		logger.Log.Info("new http server with handlers created", "host", config.All.HTTP.Host)
		err := srv.Run()
		logger.Log.Info("http server running")
		if err != nil {
			logger.Log.Warn("http starting error", "error", err)
			os.Exit(1)
		}
		ctx, shutdown := context.WithTimeout(context.Background(), consts.Timeout)
		defer shutdown()
		if err := srv.Stop(ctx); err != nil {
			log.Println("failed to stop server")
		}
	}
}
