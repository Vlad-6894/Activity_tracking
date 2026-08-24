// [ИЗМЕНЕНО]: Имя пакета изменено на core_db
package core_db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	core_config "tg-echo-bot/golang_school/internal/core/config"
)

type Pool interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Close()
}

func New(ctx context.Context, cfg core_config.PostgresConfig) (*pgxpool.Pool, error) {
	// [ИЗМЕНЕНО]: Из DSN убран timeout=...
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse postgres pool config: %w", err)
	}

	// [ИЗМЕНЕНО]: Таймаут теперь применяется через context.WithTimeout
	connCtx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(connCtx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}

	if err := pool.Ping(connCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres pool: %w", err)
	}

	return pool, nil
}
