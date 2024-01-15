package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"hh-go-bot/internal/config"
)

type Client interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func New() (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s", config.All.PostgreSQL.DBName, config.All.PostgreSQL.UserName,
		config.All.PostgreSQL.Pwd, config.All.PostgreSQL.Host, config.All.PostgreSQL.Port, config.All.PostgreSQL.DBName)
	pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return
	}
	return
}
