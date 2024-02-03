package usecase

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
)

type ResumeUsecase struct {
	request Requester
	resumes Resumes
}

func NewResumeUsecase(s RequestUsecase) ResumeUsecase {
	return ResumeUsecase{
		request: s,
		resumes: ResumeUsecase{},
	}
}

// GetResume нужен для получения id моего резюме для поиска подходящих вакансий
func (r ResumeUsecase) GetResume(ctx context.Context) (entity.Resume, error) {
	resumes := entity.NewResume()
	go r.request.Request(ctx, consts.ResumeLink)
	var body []byte
	err := json.Unmarshal(body, &resumes)
	if err != nil {
		return entity.Resume{}, fmt.Errorf("respond deserialization error: %w", err)
	}
	return resumes, nil
}
