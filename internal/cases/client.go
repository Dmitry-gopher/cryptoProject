package cases

import (
	"context"
	"cryptoProject/internal/entities"
)

type Client interface {
	GetCurrentRates(ctx context.Context, titles []string) ([]entities.Coin, error)
}
