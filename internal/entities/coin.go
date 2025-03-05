package entities

import (
	"github.com/pkg/errors"
	"time"
)

type Coin struct {
	Title       string
	CurrentRate float64
	Timestamp   time.Time
}

func NewCoin(title string, rate float64, timestamp time.Time) (*Coin, error) {
	if title == "" {
		return nil, errors.Wrap(ErrInvalidParameter, "Empty title")
	}
	if rate < 0 {
		return nil, errors.Wrap(ErrInvalidParameter, "Negative rate")
	}
	return &Coin{
		Title:       title,
		CurrentRate: rate,
		Timestamp:   timestamp,
	}, nil
}
