package cryptocompare

import (
	"context"
	"cryptoProject/internal/entities"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type CryptoCompareClient struct {
	httpClient *http.Client
}

func NewCryptoCompareClient() *CryptoCompareClient {
	return &CryptoCompareClient{
		httpClient: &http.Client{},
	}
}

func (c *CryptoCompareClient) GetCurrentRates(ctx context.Context, titles []string) ([]entities.Coin, error) {
	if len(titles) == 0 {
		return nil, errors.New("no titles")
	}

	var coins []entities.Coin

	for _, title := range titles {

		url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=USD", title)

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, errors.Wrapf(entities.Err, "failed to create request for %s", title)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, errors.Wrapf(entities.Err, "failed to send request for %s", title)
		}
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Wrapf(entities.Err, "unexpected status code for %s: %d", title, resp.StatusCode)
		}

		var data map[string]float64
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, errors.Wrapf(entities.Err, "failed to decode response for %s", title)
		}

		if err := resp.Body.Close(); err != nil {
			return nil, errors.Wrapf(entities.Err, "failed to close response body for %s", title)
		}

		rate, ok := data["USD"]
		if !ok {
			return nil, errors.Wrapf(entities.Err, "no USD in response for %s", title)
		}

		coin := entities.Coin{
			Title:       title,
			CurrentRate: rate,
			Timestamp:   time.Now(),
		}

		coins = append(coins, coin)
	}

	return coins, nil
}
