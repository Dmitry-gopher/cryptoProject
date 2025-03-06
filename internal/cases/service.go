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

//метод вызыввает метод со списком монет который получает из бд. список мб пустым
//далее вызывает метод парсинга и туда передает этот список (если пустой список - дефолтные тайтлы, иначе не пустой)
//далее метод стор
