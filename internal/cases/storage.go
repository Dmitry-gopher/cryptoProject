package cases

import (
	"context"
	"cryptoProject/internal/entities"
)

type Storage interface {
	Store(ctx context.Context, coins []entities.Coin) error
	Get(ctx context.Context, titles []string, opts ...Option) ([]entities.Coin, error)
}
