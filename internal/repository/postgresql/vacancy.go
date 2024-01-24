package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"hh-go-bot/internal/entity"
	"hh-go-bot/pkg/logger"
)

type VacancyRepo struct {
	db Client
}

func NewVacancyRepository(db Client) VacancyRepo {
	return VacancyRepo{db: db}
}

func (r VacancyRepo) Create(ctx context.Context, vacancies entity.Vacancies) error {
	for _, v := range vacancies.Items {
		query := `INSERT INTO public.vacancies (icon, id, name, relations,"employer.name", published_at, created_at, archived, alternate_url, "experience.id", "experience.name") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11)`
		if err := r.db.QueryRow(ctx, query, v.Icon, v.Id, v.Name, v.Relations, v.Employer.Name, v.PublishedAt, v.CreatedAt, v.Archived, v.AlternateUrl, v.Experience.ID, v.Experience.Name).Scan(v.Id); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				fmt.Printf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			}
		}
		logger.Log.Debug("vacancy added to db", "id", v.Id)
	}
	return nil
}

func (r VacancyRepo) Update(ctx context.Context, v entity.Vacancies) error {
	return nil
}
func (r VacancyRepo) One(ctx context.Context, id string) (entity.Vacancies, error) {
	return entity.Vacancies{}, nil
}
func (r VacancyRepo) AllResponded(ctx context.Context) (entity.Vacancies, error) {
	return entity.Vacancies{}, nil

}
func (r VacancyRepo) Delete(ctx context.Context, id string) error {
	return nil
}
