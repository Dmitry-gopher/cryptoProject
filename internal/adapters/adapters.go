package adapters

// TODO подумать над обработкой ошибок, спросить про дефолтный кейс в Get
import (
	"context"
	"cryptoProject/internal/cases"
	"cryptoProject/internal/entities"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type StorageImplementation struct {
	pool *pgxpool.Pool
}

func NewStorage(ctx context.Context, connString string) (*StorageImplementation, error) {
	pool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, errors.Wrap(entities.ErrStorageGetFailed, "failed to connect to database")
	}
	return &StorageImplementation{
		pool: pool,
	}, nil
}

func (s *StorageImplementation) Store(ctx context.Context, coins []entities.Coin) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return errors.Wrap(entities.ErrStorageGetFailed, "failed to store rate")
	}
	defer tx.Rollback(ctx)

	for _, coin := range coins {
		_, err := tx.Exec(ctx, "INSERT INTO coins (title, current_price, timestamp) VALUES ($1, $2, $3)",
			coin.Title, coin.CurrentPrice, coin.Timestamp)
		if err != nil {
			return errors.Wrap(entities.ErrStorageGetFailed, "failed to store rate")
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return errors.Wrap(entities.ErrStorageGetFailed, "failed to commit transaction")
	}
	return nil
}

func (s *StorageImplementation) Get(ctx context.Context, titles []string, opts ...cases.Option) ([]entities.Coin, error) {
	options := &cases.Options{}
	for _, opt := range opts {
		opt(options)
	}

	var query string
	switch options.FuncType {
	case cases.Max:
		query = "SELECT title, MAX(current_price) as current_price, timestamp FROM coins WHERE title = ANY($1) GROUP BY title"
	case cases.Min:
		query = "SELECT title, MIN(current_price) as current_price, timestamp FROM coins WHERE title = ANY($1) GROUP BY title"
	case cases.Avg:
		query = "SELECT title, AVG(current_price) as current_price, timestamp FROM coins WHERE title = ANY($1) GROUP BY title"
	default:
		query = "SELECT DISTINCT ON (title) title, current_price, timestamp FROM coins WHERE title = ANY($1) GROUP BY title, timestamp DESC"
	}

	rows, err := s.pool.Query(ctx, query, titles)
	if err != nil {
		return nil, errors.Wrap(entities.ErrStorageGetFailed, "failed to execute query")
	}

	defer rows.Close()
	var coins []entities.Coin
	for rows.Next() {
		var coin entities.Coin
		if err := rows.Scan(&coin.Title, &coin.CurrentPrice, &coin.Timestamp); err != nil {
			return nil, errors.Wrap(entities.ErrStorageGetFailed, "failed to scan row")
		}
		coins = append(coins, coin)
	}
	if rows.Err() != nil {
		return nil, errors.Wrap(rows.Err(), "row error")
	}

	return coins, nil
}
