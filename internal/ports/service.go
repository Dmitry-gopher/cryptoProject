package ports

import (
	"context"
	"cryptoProject/internal/ports/dto"
)

type Service interface {
	GetMaxRate(ctx context.Context, titles []string) (dto.Coins, error)
	GetMinRate(ctx context.Context, titles []string) (dto.Coins, error)
	GetAvgRate(ctx context.Context, titles []string) (dto.Coins, error)
	GetLastRate(ctx context.Context, titles []string) (dto.Coins, error)
}
