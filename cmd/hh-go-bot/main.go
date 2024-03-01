package main

import (
	"context"
	"flag"
	tele "gopkg.in/telebot.v3"
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
	"hh-go-bot/transport/telegrambot/bothandler"
	"os"
	"time"
)

func main() {
	redisClient := cache.NewRedisClient(config.All.Redis.Host, config.All.Redis.Pwd)
	redisService := cache.NewRedisService(redisClient)
	logger.Log.Info("redis client added", "address", config.All.Redis.Host)

	db, err := postgresql.New(config.All)
	repos := repository.NewRepositories(db, redisService)
	if err != nil {
		logger.Log.Error("repo creating error", "error", err)
		os.Exit(1)
	}

	services := service.NewServices(repos)

	f := flag.String("d", "bot", "delivery using")
	flag.Parse()
	logger.Log.Info("starting application", "mode", f)

	switch *f {
	case consts.BOT:
		pref := tele.Settings{
			Token:  config.All.Bot.Token,
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		}
		b := telegrambot.NewBotService(pref, services)
		h := bothandler.NewHandler(b)
		h.Init()
		b.Bot.Start()

	case consts.HTTP:
		handlers := handler.NewHandler(services)
		logger.Log.Info("handlers initialized")
		srv := http.NewServer(config.All.HTTP.Host, handlers.Init())
		logger.Log.Info("new http server with handlers created", "host", config.All.HTTP.Host)
		err := srv.Run()
		logger.Log.Info("http server running")
		if err != nil {
			logger.Log.Error("http starting error", "error", err)
			os.Exit(1)
		}
		ctx, shutdown := context.WithTimeout(context.Background(), consts.Timeout)
		defer shutdown()
		if err := srv.Stop(ctx); err != nil {
			logger.Log.Error("failed to stop server", err.Error())
		}
	}
}
