package telegrambot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hh-go-bot/internal/service"
	"time"
)

type BotInterface interface {
	Start() error
	SendArrayMessage(*tgbotapi.BotAPI, tgbotapi.MessageConfig, []string)
}

type BotService struct {
	botInterface BotInterface
	botapi       *tgbotapi.BotAPI
	services     service.Services
}

func NewBot(token string, services service.Services) (*BotService, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &BotService{
		botapi:       bot,
		services:     services,
		botInterface: BotService{},
	}, nil
}

func (b BotService) Start() (err error) {
	b.botapi.Debug = true
	u := tgbotapi.NewUpdate(0)
	updates := b.botapi.GetUpdatesChan(u)
	for update := range updates {
		ch := make(chan []string)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.DisableWebPagePreview = true
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		switch update.Message.Command() {

		case "similar":
			go b.services.Vacancier.Similar(ctx, ch)

		case "jobs":
			go b.services.Vacancier.All(ctx, ch)

		case "resume":
			go b.services.Resumes.MyResume(ctx, ch)

		default:
			go func() {
				ch <- append([]string{}, "I don't know that command")
			}()
		}
		select {
		case <-ctx.Done():
			msg.Text = "20 sec timeout"
		case text := <-ch:
			b.botInterface.SendArrayMessage(b.botapi, msg, text)
		}
		close(ch)
	}
	return
}

func (b BotService) SendArrayMessage(bapi *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, m []string) {
	for _, v := range m {
		msg.Text = v
		if _, err := bapi.Send(msg); err != nil {
			fmt.Println("error")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
