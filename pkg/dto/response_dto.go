package dto

import "time"

type CoinDTO struct {
	Title       string    `json:"title"`
	CurrentRate float64   `json:"current_rate"`
	Timestamp   time.Time `json:"timestamp"`
}

type CoinsDTO []CoinDTO
