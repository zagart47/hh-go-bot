package tool

import (
	"hh-go-bot/internal/model/job"
	"sort"
)

type Converter interface {
	MapToSlice(map[string]job.Vacancy) job.VacancyList
}

type Service struct {
	tool Converter
}

func (t Service) MapToSlice(m map[string]job.Vacancy) job.VacancyList {
	vacancies := job.VacancyList{}
	keys := make([]string, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		vacancies.Items = append(vacancies.Items, m[k])
	}
	return vacancies
}
