package public

import (
	"context"
	"cryptoProject/internal/entities"
)

type Service interface {
	GetMaxRate(ctx context.Context, titles []string) ([]entities.Coin, error)
	GetMinRate(ctx context.Context, titles []string) ([]entities.Coin, error)
	GetAvgRate(ctx context.Context, titles []string) ([]entities.Coin, error)
	GetLastRate(ctx context.Context, titles []string) ([]entities.Coin, error)
}
