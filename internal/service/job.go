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
	"strings"
)

type Job interface {
	All() entity.VacancyList
	Message(entity.VacancyList) []string
}

type jobService struct {
	job       Job
	converter Converter
}

func (s jobService) Message(vacancyList entity.VacancyList) []string {
	var message string
	var messageList []string
	var count int
	for _, v := range vacancyList.Items {
		message = fmt.Sprintf("%s\n%c %s | %s - %s", message, v.Applied, v.Employer.Name, v.Name, v.AlternateUrl)
		count++
		if count == 40 { // 40 вакансий - примерное оптимальное количество для непревышения лимита символов в одном сообщении (4096)
			messageList = append(messageList, message)
			count = 0
			message = ""
		}
	}
	if message != "" {
		messageList = append(messageList, message)
		return messageList
	}
	return nil
}

func NewService() jobService {
	return jobService{
		jobService{},
		toolService{},
	}
}

func (s jobService) All() entity.VacancyList {
	listMap := make(map[string]entity.Vacancy)
	list := entity.NewVacancyList()
	for i := 0; ; i++ {
		url := fmt.Sprintf("https://api.hh.ru/vacancies?text=golang&area=113&id=publication_time&page=%d", i)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
		}

		body, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Ошибка при чтении ответа:", err)
		}

		err = json.Unmarshal(body, &list)
		if err != nil {
			fmt.Println("Ошибка при десериализации ответа:", err)
		}
		for _, v := range list.Items {
			if strings.Contains(strings.ToLower(v.Name), "go") {
				if FindAppliedVacancies(v.Id) {
					v.Applied = '\u2705'
				}
				listMap[v.PublishedAt+v.Id] = v
			}
		}
		if list.Pages == i {
			break
		}
	}
	return s.converter.MapToSlice(listMap)
}

func (s jobService) Similar() entity.VacancyList {
	listMap := make(map[string]entity.Vacancy)
	var list entity.VacancyList
	var link string
	cfg, err := config.New()
	if err != nil {
		log.Print(err)
	}
	for i := 0; ; i++ {
		buffer := bytes.NewBuffer([]byte(`{"key": "value"}`))
		link = fmt.Sprintf("https://api.hh.ru/resumes/%s/similar_vacancies?area=113&id=relevance&page=%d", cfg.ResumeID, i)
		request, err := http.NewRequest("GET", link, buffer)
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
		for _, v := range list.Items {
			if strings.Contains(strings.ToLower(v.Name), "go") {
				listMap[v.PublishedAt+v.Id] = v
			}
		}
		if list.Pages == i {
			break
		}
	}
	return s.converter.MapToSlice(listMap)
}
