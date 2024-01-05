package service

import (
	"hh-go-bot/internal/entity"
	"sort"
)

type Converter interface {
	MapToSlice(map[string]entity.Vacancy) entity.VacancyList
}

type toolService struct {
	tool Converter
}

func (t toolService) MapToSlice(m map[string]entity.Vacancy) entity.VacancyList {
	vacancies := entity.VacancyList{}
	keys := make([]string, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys) // sort vacancies by id

	for _, k := range keys {
		vacancies.Items = append(vacancies.Items, m[k])
	}
	return vacancies
}
