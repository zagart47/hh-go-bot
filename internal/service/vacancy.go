package service

import (
	"context"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/usecase"
	"hh-go-bot/pkg/logger"
)

type VacancyService struct {
	vacancy Vacancy
	usecase usecase.Usecases
}

func (s VacancyService) One(ctx context.Context, id string) (entity.Vacancy, error) {
	m, err := s.usecase.Vacancy.GetOne(ctx, id)
	if err != nil {
		logger.Log.Error("vacancy getting error", err.Error())
	}
	return m, nil
}

func NewVacancyService(usecase usecase.Usecases) VacancyService {
	return VacancyService{
		vacancy: VacancyService{},
		usecase: usecase,
	}
}

func (s VacancyService) All(ctx context.Context) (entity.Vacancies, error) {
	m, err := s.usecase.Vacancy.GetAll(ctx, consts.AllVacanciesLink)
	if err != nil {
		return entity.Vacancies{}, err
	}
	v := s.usecase.Sorter.Sort(ctx, m)
	return v, nil
}

func (s VacancyService) Similar(ctx context.Context) (entity.Vacancies, error) {
	m, err := s.usecase.Vacancy.GetAll(ctx, consts.SimilarVacanciesLink)
	if err != nil {
		return entity.Vacancies{}, err
	}
	v := s.usecase.Sorter.Sort(ctx, m)
	return v, nil
}
