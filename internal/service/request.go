package service

import (
	"bytes"
	"context"
	"fmt"
	"hh-go-bot/internal/config"
	"io"
	"net/http"
)

type RequestService struct {
	request Requester
}

func NewRequestService() RequestService {
	return RequestService{}
}

// Do отправляет запросы с bearer токеном
func (r RequestService) doRequest(_ context.Context, link string, ch chan []byte) {

	buffer := bytes.NewBuffer([]byte(`{"key": "value"}`))
	request, err := http.NewRequest(http.MethodGet, link, buffer)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.All.Api.Bearer))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Ошибка при отправке запроса: %v\n", err)
	}

	raw, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
	}

	defer response.Body.Close()
	ch <- raw
}
