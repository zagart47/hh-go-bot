package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/entity"
	"io"
	"log"
	"net/http"
	"os"
)

type Vacancies interface {
	FindAppliedVacancies() entity.VacancyList
}

func FindAppliedVacancies(id string) bool {
	list := entity.NewAppliedVacancy()
	cfg, err := config.New()
	if err != nil {
		log.Print(err)
	}
	buffer := bytes.NewBuffer([]byte(`{"key": "value"}`))
	link := fmt.Sprintf("https://api.hh.ru/vacancies/%s/suitable_resumes", id)
	request, err := http.NewRequest("GET", link, buffer)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Authorization", "Bearer "+cfg.Bearer)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Ошибка при отправке запроса: %v\n", err)
		os.Exit(1)
	}
	respBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
	}
	err = json.Unmarshal(respBody, &list)
	if err != nil {
		fmt.Println("Ошибка при десериализации ответа:", err)
	}
	if list.Overall.AlreadyApplied == 1 {
		return true
	}
	return false
}
