package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"hh-go-bot/internal/entity"
)

type Vacancies interface {
	Create(context.Context, entity.Vacancies) error
	One(context.Context) (entity.Vacancies, error)
	AllResponded(context.Context) (entity.Vacancies, error)
	Delete(context.Context, string) error
}

type Repositories struct {
	Vacancies Vacancies
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		Vacancies: NewVacancyRepository(db),
	}
}
