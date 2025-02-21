package entities

import "time"

type Coin struct {
	Title        string
	CurrentPrice float64
	Timestamp    time.Time
}

func NewCoin(title string, price float64, timestamp time.Time) *Coin {
	return &Coin{
		Title:        title,
		CurrentPrice: price,
		Timestamp:    timestamp,
	}
}
