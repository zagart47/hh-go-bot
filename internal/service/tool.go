package service

import (
	"hh-go-bot/internal/model"
	"sort"
)

type Converter interface {
	MapToSlice(map[string]model.Vacancy) model.VacancyList
}

type toolService struct {
	tool Converter
}

func (t toolService) MapToSlice(m map[string]model.Vacancy) model.VacancyList {
	vacancies := model.VacancyList{}
	keys := make([]string, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys) // sorting by vacancies id

	for _, k := range keys {
		vacancies.Items = append(vacancies.Items, m[k])
	}
	return vacancies
}
