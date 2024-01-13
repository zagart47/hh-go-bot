package service

import (
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
	"sort"
)

const (
	between1and3 = "between1And3"
	between3and6 = "between3And6"
	moreThan6    = "moreThan6"
)

type ConverterService struct {
	converter Converter
}

func NewConverterService() ConverterService {
	return ConverterService{}
}

func (c ConverterService) convert(m map[string]entity.Vacancy) entity.Vacancies {
	vacancies := entity.NewVacancies()
	keys := make([]string, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys) // sort vacancies by publication date+id
	for _, k := range keys {
		if m[k].Experience.ID == consts.NoExperience {
			vacancies.Items = append(vacancies.Items, m[k])
			delete(m, k)
		}
	}
	for _, k := range keys {
		if m[k].Experience.ID == between1and3 {
			vacancies.Items = append(vacancies.Items, m[k])
			delete(m, k)
		}
	}
	for _, k := range keys {
		if m[k].Experience.ID == between3and6 {
			vacancies.Items = append(vacancies.Items, m[k])
			delete(m, k)
		}
	}
	for _, k := range keys {
		if m[k].Experience.ID == moreThan6 {
			vacancies.Items = append(vacancies.Items, m[k])
			delete(m, k)
		}
	}
	return vacancies
}
