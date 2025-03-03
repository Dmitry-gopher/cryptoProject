package cases

import (
	"context"
	"cryptoProject/internal/entities"
)

type CryptoClient interface {
	GetCurrentRates(ctx context.Context, titles []string) ([]entities.Coin, error)
}
