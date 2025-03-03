package entities

import (
	"github.com/pkg/errors"
	"time"
)

type Coin struct {
	Title        string
	CurrentPrice float64
	Timestamp    time.Time
}

func NewCoin(title string, price float64, timestamp time.Time) (*Coin, error) {
	if title == "" {
		return nil, errors.Wrap(ErrInvalidParameter, "Empty title")
	}
	if price < 0 {
		return nil, errors.Wrap(ErrInvalidParameter, "Negative rate")
	}
	return &Coin{
		Title:        title,
		CurrentPrice: price,
		Timestamp:    timestamp,
	}, nil
}
