package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hh-go-bot/internal/config"
	"hh-go-bot/internal/model"
	"io"
	"net/http"
	"os"
	"strings"
)

type Job interface {
	All() model.VacancyList
	Message(model.VacancyList) []string
}

type jobService struct {
	job  Job
	tool Converter
}

func (js jobService) Message(vacancyList model.VacancyList) []string {
	var message string
	var messageList []string
	var count int
	for _, v := range vacancyList.Items {
		message = fmt.Sprintf("%s\n%s | %s - %s", message, v.Employer.Name, v.Name, v.AlternateUrl)
		count++
		if count == 40 { // примерно 40 вакансий - оптимальное количество вакансий для непревышения лимита символов в одном сообщении (4096)
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

func (js jobService) MapToSlice(m map[string]model.Vacancy) model.VacancyList {
	return js.tool.MapToSlice(m)
}

func (js jobService) All() model.VacancyList {
	listMap := make(map[string]model.Vacancy)
	var list model.VacancyList
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
				listMap[v.PublishedAt+v.Id] = v
			}
		}
		if list.Pages == i {
			break
		}
	}
	return js.MapToSlice(listMap)
}

func (js jobService) Similar() model.VacancyList {
	listMap := make(map[string]model.Vacancy)
	var list model.VacancyList
	var link string
	for i := 0; ; i++ {
		buffer := bytes.NewBuffer([]byte(`{"key": "value"}`))
		link = fmt.Sprintf("https://api.hh.ru/resumes/e8c5ba8eff0cac22500039ed1f446166626974/similar_vacancies?area=113&id=relevance&page=%d", i)
		request, err := http.NewRequest("GET", link, buffer)
		request.Header.Set("Authorization", "Bearer "+config.New().Bearer)
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
	return js.MapToSlice(listMap)
}
