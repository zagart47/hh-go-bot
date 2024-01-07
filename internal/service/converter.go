package service

import (
	"context"
	"hh-go-bot/internal/entity"
	"sort"
)

type ConverterService struct {
	converter Converter
}

func NewConverterService() ConverterService {
	return ConverterService{}
}

const (
	Between1and3 = "between1And3"
	Between3and6 = "between3And6"
	NoExperience = "noExperience"
	MoreThan6    = "moreThan6"
)

func (c ConverterService) Convert(ctx context.Context, m map[string]entity.Vacancy) entity.Vacancies {
	vacancies := entity.NewVacancies()
	keys := make([]string, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys) // sort vacancies by id

	for _, k := range keys {
		if m[k].Experience.ID == NoExperience {
			vacancies.Items = append(vacancies.Items, m[k])
		}
	}
	for _, k := range keys {
		if m[k].Experience.ID == Between1and3 {
			vacancies.Items = append(vacancies.Items, m[k])
		}
	}
	for _, k := range keys {
		if m[k].Experience.ID == Between3and6 {
			vacancies.Items = append(vacancies.Items, m[k])
		}
	}
	for _, k := range keys {
		if m[k].Experience.ID == MoreThan6 {
			vacancies.Items = append(vacancies.Items, m[k])
		}
	}
	return vacancies
}
