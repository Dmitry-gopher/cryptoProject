package storage

import (
	"context"
	"cryptoProject/internal/entities"
)

type CoinData struct { // Data Transfer Object (DTO) для ответа
	CurrentPrice float64
	HourlyChange *float64 // |
	DayMin       *float64 // | nil, если не запрошено
	DayMax       *float64 // |
}

type QueryOption func(*queryOptions)

type queryOptions struct {
	hourlyChange bool
	dailyChange  bool
}

func WithHourlyChange() QueryOption {
	return func(o *queryOptions) { o.hourlyChange = true }
}

func WithDailyChange() QueryOption {
	return func(o *queryOptions) { o.dailyChange = true }
}

type CurrencyRepository interface {
	SaveRate(ctx context.Context, c *entities.Coin) error
	Get(ctx context.Context, CoinTitle string, opts ...QueryOption) (*CoinData, error)
}

/*
Пример запроса: Get("BTC", WithHourlyChange(), WithDailyChange())
- тогда SELECT будет объёмным и запрашивать цену
и подсчёт всех нужных изменений;
простой запрос Get("BTC") вернёт только цену, без подсчётов.
Тогда сложность запросов будет на уровне БД, все расчёты там же,
один запрос вместо нескольких,
приложение сохраняется простым и тестируемым (наверно...)
*/
