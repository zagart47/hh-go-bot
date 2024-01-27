package usecase

import (
	"context"
	"hh-go-bot/internal/consts"
	"hh-go-bot/internal/entity"
	"sort"
)

const (
	between1and3 = "between1And3"
	between3and6 = "between3And6"
	moreThan6    = "moreThan6"
)

type SortUsecase struct {
	Sorter Sorter
}

func NewSortUsecase() SortUsecase {
	return SortUsecase{}
}

// Sort необходим для сортировки вакансий по опыту
func (p SortUsecase) Sort(ctx context.Context, m map[string]entity.Vacancy) entity.Vacancies {
	v := entity.NewVacancies()
	keys := make([]string, 0, len(m))
	for _, k := range m {
		keys = append(keys, k.Id)
	}
	sort.Strings(keys) // pick vacancies by publication date+id
	for _, k := range keys {
		if m[k].Exp.ID == consts.NoExp {
			v.Items = append(v.Items, m[k])
			delete(m, k)
		}
	}
	for _, k := range keys {
		if m[k].Exp.ID == between1and3 {
			v.Items = append(v.Items, m[k])
			delete(m, k)
		}
	}
	for _, k := range keys {
		if m[k].Exp.ID == between3and6 {
			v.Items = append(v.Items, m[k])
			delete(m, k)
		}
	}
	for _, k := range keys {
		if m[k].Exp.ID == moreThan6 {
			v.Items = append(v.Items, m[k])
			delete(m, k)
		}
	}
	return v
}
