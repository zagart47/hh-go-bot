package postgresql

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"hh-go-bot/internal/entity"
	"hh-go-bot/pkg/logger"
	"time"
)

type VacancyRepo struct {
	db Client
}

func NewVacancyRepository(db Client) VacancyRepo {
	return VacancyRepo{db: db}
}

func (r VacancyRepo) Create(ctx context.Context, vacancies entity.Vacancies) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	for _, v := range vacancies.Items {
		query := `
		INSERT INTO public.vacancies (icon, id, name, relations,"employer.name", published_at, created_at, archived, alternate_url, "experience.id", "experience.name")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING (id)`
		if err := r.db.QueryRow(ctx, query, v.Icon, v.Id, v.Name, v.Relations, v.Employer.Name, v.PublishedAt, v.CreatedAt, v.Archived, v.AlternateUrl, v.Exp.ID, v.Exp.Name).Scan(&v.Id); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				logger.Log.Warn("sql error:", pgErr.Message, "detail:", pgErr.Detail, "where:", pgErr.Where)
				//fmt.Printf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Prepare, pgErr.Detail, pgErr.Where)
			} else if err != nil {
				return err
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
