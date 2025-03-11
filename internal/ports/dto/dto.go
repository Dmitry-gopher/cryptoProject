package dto

type Coin struct {
	Title       string  `json:"title"`
	CurrentRate float64 `json:"current_rate"`
	Timestamp   string  `json:"timestamp"`
}

type Coins []Coin
