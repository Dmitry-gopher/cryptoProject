package usecases

import (
	"context"
	"cryptoProject/internal/entities"
)

type CryptoAPI interface {
	GetCurrentRates(ctx context.Context, coinsTitles []string) ([]entities.Coin, error)
}
