package service

import (
	"bytes"
	"context"
	"fmt"
	"hh-go-bot/internal/config"
	"io"
	"net/http"
	"os"
)

type RequestService struct {
	request Requester
}

func NewRequestService() RequestService {
	return RequestService{}
}

// Request отправляет запросы с bearer токеном
func (r RequestService) Request(ctx context.Context, link string) []byte {
	cfg, err := config.All()
	if err != nil {
		fmt.Println(err)
	}
	buffer := bytes.NewBuffer([]byte(`{"key": "value"}`))
	request, err := http.NewRequest("GET", link, buffer)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.Api.Bearer))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Ошибка при отправке запроса: %v\n", err)
		os.Exit(1)
	}
	raw, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
	}
	return raw
}
