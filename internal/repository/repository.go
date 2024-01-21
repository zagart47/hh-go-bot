package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"hh-go-bot/internal/entity"
	"hh-go-bot/internal/repository/cache"
	"hh-go-bot/internal/repository/postgresql"
)

type Vacancies interface {
	Create(context.Context, entity.Vacancies) error
	Update(context.Context, entity.Vacancies) error
	One(context.Context, string) (entity.Vacancies, error)
	AllResponded(context.Context) (entity.Vacancies, error)
	Delete(context.Context, string) error
}

type Repositories struct {
	Vacancies Vacancies
	Redis     cache.RedisService
}

func NewRepositories(db *pgxpool.Pool, redis cache.RedisService) Repositories {
	return Repositories{
		Vacancies: postgresql.NewVacancyRepository(db),
		Redis:     redis,
	}
}
