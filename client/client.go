package client

type CryptoAPI interface {
	FetchCurrentRates(currencyIDs []string) (map[string]float64, error)
}
