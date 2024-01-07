package service

import (
	"context"
	"encoding/json"
	"fmt"
	"hh-go-bot/internal/entity"
)

type ResumeService struct {
	request Requester
	resumes Resumes
}

func NewResumeService(service RequestService) ResumeService {
	return ResumeService{
		request: service,
		resumes: ResumeService{},
	}
}

// MyResume нужен для получения id моего резюме для поиска подходящих вакансий
func (r ResumeService) MyResume(ctx context.Context, ch chan []string) {
	resumes := entity.NewResume()
	link := "https://api.hh.ru/resumes/mine"
	body := r.request.Request(ctx, link)
	err := json.Unmarshal(body, &resumes)
	if err != nil {
		fmt.Println("Ошибка при десериализации ответа:", err)
	}
	var resumeID []string
	for _, v := range resumes.Items {
		resumeID = append(resumeID, v.ID)
	}
	ch <- resumeID
}
