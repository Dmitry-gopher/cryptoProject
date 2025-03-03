package cases

import (
	"context"
	"cryptoProject/internal/entities"
)

// TODO
pgx pool 4.0 (структура, конструктор, пул соединений в конструкторе), методы стор и гет. Гпт
может помочь. Попробовать запросы в sql-online (примерно 6 записей, 2 монеты, последняя запись не макс,
	мин, ср) -
	скормить гпт, попросить селект запрос и проверить его. Это для проверки работы запросов. Всё это в отдельном слое
адаптеров.

type Storage interface {
	Store(ctx context.Context, coins []entities.Coin) error
	Get(ctx context.Context, titles []string, opts ...Option) ([]entities.Coin, error)
}