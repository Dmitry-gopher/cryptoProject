package core

import "time"

type Currency struct {
	Name         string       // Полное название монеты
	Symbol       string       // Тикер монеты
	CurrentPrice float64      // Текущая цена в USD
	PriceHistory []PricePoint // Исторические данные для расчётов
}
type PricePoint struct {
	Timestamp time.Time
	Price     float64
}
