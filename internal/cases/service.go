package cases

import (
	"context"
	"cryptoProject/internal/entities"
	"github.com/pkg/errors"
)

type Service struct {
	storage Storage
	client  Client
}

func NewService(storage Storage, client Client) (*Service, error) {
	if storage == nil {
		return nil, errors.Wrap(entities.ErrInvalidParameter, "storage can't be nil")
	}
	if client == nil {
		return nil, errors.Wrap(entities.ErrInvalidParameter, "client can't be nil")
	}
	return &Service{
		storage: storage,
		client:  client,
	}, nil
}

type AggFunc int

const (
	_ AggFunc = iota
	Max
	Min
	Avg
)

type Options struct {
	FuncType AggFunc
}

type Option func(options *Options)

func (a AggFunc) String() string {
	return [...]string{"", "MAX", "MIN", "AVG"}[a]
}

func WithMax() Option {
	return func(options *Options) {
		options.FuncType = Max
	}
}

func (s *Service) GetMaxRate(ctx context.Context, titles []string) ([]entities.Coin, error) {
	coins, err := s.storage.Get(ctx, titles, WithMax())
	if err != nil {
		return nil, errors.Wrap(entities.ErrStorageGetFailed, "Failed to get max rates")
	}
	return coins, nil
}

func WithMin() Option {
	return func(options *Options) {
		options.FuncType = Min
	}
}

func (s *Service) GetMinRate(ctx context.Context, titles []string) ([]entities.Coin, error) {
	coins, err := s.storage.Get(ctx, titles, WithMin())
	if err != nil {
		return nil, errors.Wrap(entities.ErrStorageGetFailed, "Failed to get min rates")
	}
	return coins, nil
}

func WithAvg() Option {
	return func(options *Options) {
		options.FuncType = Avg
	}
}

func (s *Service) GetAvgRate(ctx context.Context, titles []string) ([]entities.Coin, error) {
	coins, err := s.storage.Get(ctx, titles, WithAvg())
	if err != nil {
		return nil, errors.Wrap(entities.ErrStorageGetFailed, "Failed to get average rates")
	}
	return coins, nil
}

func (s *Service) GetLastRate(ctx context.Context, titles []string) ([]entities.Coin, error) {
	coins, err := s.storage.Get(ctx, titles)
	if err != nil {
		return nil, errors.Wrap(entities.ErrStorageGetFailed, "Failed to get last rates")
	}
	return coins, nil
}

func (s *Service) UpdateCoinRates(ctx context.Context) error {
	titles, err := s.storage.GetUniqueTitles(ctx)
	if err != nil {
		return errors.Wrap(entities.ErrStorageGetFailed, "Failed to get unique titles")
	}

	coins, err := s.client.GetCurrentRates(ctx, titles)
	if err != nil {
		return errors.Wrap(entities.ErrStorageGetFailed, "Failed to get current rates")
	}

	if err := s.storage.Store(ctx, coins); err != nil {
		return errors.Wrap(err, "Failed to store coins")
	}

	return nil
}

//слой портов (будет интерфейс сервиса). структура сервера, конструктор, метод запуска Run
//внутри будут вызываться методы сервиса, кроме последнего
//на уровне портов не можем использовать entities - нужен отдельный пакет - dto (2 шт: coin и слайс coin)
//для них конструктор не нужен
//пакеты chi.mux - для роутинга, http-константы ошибок для обработки их
//
//swagger - для документации. собрать application - отдельный слой app (метод Run - запуск всего)
//чтобы парсинг был в фоне - cron джоба (пакет cron)
