package cryptocompare

import (
	"context"
	"cryptoProject/internal/entities"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CryptoCompareClient struct {
	// поле со списком дефолтных монет
	httpClient *http.Client
	baseURL    string
}

func NewCryptoCompareClient() *CryptoCompareClient {
	return &CryptoCompareClient{
		httpClient: &http.Client{},
		baseURL:    "https://min-api.cryptocompare.com/data/pricemulti",
	}
}

func (c *CryptoCompareClient) GetCurrentRates(ctx context.Context, titles []string) ([]entities.Coin, error) {
	if len(titles) == 0 {
		return nil, errors.Wrap(entities.ErrInvalidParameter, "no titles")
	} // возвращать не нил, а список дефолтных монет

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url")
	}
	q := u.Query()
	q.Set("fsyms", strings.Join(titles, ","))
	q.Set("tsyms", "USD")
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(err, "unexpected status code: %d", resp.StatusCode)
	}

	var data map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	var coins []entities.Coin
	for title, rates := range data {
		rate, ok := rates["USD"]
		if !ok {
			return nil, errors.Wrapf(entities.ErrInvalidParameter, "Rate in USD isn't available for %s", title)
		}
		coins = append(coins, entities.Coin{
			Title:       title,
			CurrentRate: rate,
			Timestamp:   time.Now(),
		})
	}

	return coins, nil
}
