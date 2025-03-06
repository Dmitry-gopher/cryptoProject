package postgres

import (
	"context"
	"cryptoProject/internal/cases"
	"cryptoProject/internal/entities"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type PgStorage struct {
	pool *pgxpool.Pool
}

func NewPgStorage(ctx context.Context, connString string) (*PgStorage, error) {
	pool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, errors.Wrap(entities.ErrStorage, "failed to connect to storage")
	}
	return &PgStorage{
		pool: pool,
	}, nil
}

func (s *PgStorage) Store(ctx context.Context, coins []entities.Coin) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return errors.Wrap(entities.ErrStorage, "failed to begin transaction")
	}

	for _, coin := range coins {
		_, err := tx.Exec(ctx, "INSERT INTO coins (title, current_price, timestamp) VALUES ($1, $2, $3);",
			coin.Title, coin.CurrentRate, coin.Timestamp)
		if err != nil {
			return errors.Wrap(entities.ErrStorage, "failed to store rate into storage")
		}
	}
	if err := tx.Commit(ctx); err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return errors.Wrap(entities.ErrStorage, "failed to rollback transaction")
		}
		return errors.Wrap(entities.ErrStorage, "failed to commit transaction")
	}
	return nil
}

func (s *PgStorage) Get(ctx context.Context, titles []string, opts ...cases.Option) ([]entities.Coin, error) {
	options := &cases.Options{}
	for _, opt := range opts {
		opt(options)
	}

	var query string
	switch options.FuncType {
	case cases.Max:
		query = "SELECT title, MAX(current_price) as max_price FROM coins WHERE title = ANY($1) GROUP BY title;"
	case cases.Min:
		query = "SELECT title, MIN(current_price) as min_price FROM coins WHERE title = ANY($1) GROUP BY title;"
	case cases.Avg:
		query = "SELECT title, AVG(current_price) as avg_price FROM coins WHERE title = ANY($1) GROUP BY title;"
	default:
		query = "SELECT DISTINCT ON (title) title, current_price, timestamp FROM coins WHERE title = ANY($1) ORDER BY title, timestamp DESC;"
	}

	rows, err := s.pool.Query(ctx, query, titles)
	if err != nil {
		return nil, errors.Wrap(entities.ErrStorageGetFailed, "failed to execute query")
	}
	defer rows.Close()

	var coins []entities.Coin
	for rows.Next() {
		var coin entities.Coin
		err = rows.Scan(&coin.Title, &coin.CurrentRate, &coin.Timestamp)
		if err != nil {
			return nil, errors.Wrap(entities.ErrStorageGetFailed, "failed to scan row")
		}
		coins = append(coins, coin)
	}
	if rows.Err() != nil {
		return nil, errors.Wrap(rows.Err(), "some row error")
	}

	return coins, nil
}
