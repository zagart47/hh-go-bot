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

func New(c config.Cfg) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s", c.PostgreSQL.DBName, c.PostgreSQL.UserName,
		c.PostgreSQL.Pwd, c.PostgreSQL.Host, c.PostgreSQL.Port, c.PostgreSQL.DBName)
	pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return
	}
	return
}
