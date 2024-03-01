package bothandler

import (
	"context"
	tele "gopkg.in/telebot.v3"
	"hh-go-bot/internal/entity"
)

func (b *Bot) Resume(c tele.Context) error {
	resume, err := b.Services.Resume.Get(context.Background())
	if resume.Items == nil {
		return c.Send("У тебя нет видимых резюме")
	}
	if err != nil {
		return err
	}
	for _, r := range ResumeMessage(resume) {
		err = c.Send(r, &tele.SendOptions{DisableWebPagePreview: true})
		if err != nil {
			return err
		}
	}
	return nil
}

func ResumeMessage(r entity.Resume) []string {
	var text []string
	for _, v := range r.Items {
		text = append(text, v.ID)
	}
	return text
}
