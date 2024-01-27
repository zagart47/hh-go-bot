package service

import (
	"context"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/usecase"
)

type VacancyService struct {
	vacancy Vacancy
	usecase usecase.Usecases
}

func NewVacancyService(usecase usecase.Usecases) VacancyService {
	return VacancyService{
		vacancy: VacancyService{},
		usecase: usecase,
	}
}

func (s VacancyService) All(ctx context.Context) (entity.Vacancies, error) {
	m, err := s.usecase.Vacancy.Get(ctx, consts.AllVacanciesLink)
	if err != nil {
		return entity.Vacancies{}, err
	}
	v := s.usecase.Sorter.Sort(ctx, m)
	return v, nil
}

func (s VacancyService) Similar(ctx context.Context) (entity.Vacancies, error) {
	m, err := s.usecase.Vacancy.Get(ctx, consts.SimilarVacanciesLink)
	if err != nil {
		return entity.Vacancies{}, err
	}
	v := s.usecase.Sorter.Sort(ctx, m)
	return v, nil
}
