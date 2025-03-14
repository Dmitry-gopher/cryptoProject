package app

import (
	"context"
	"cryptoProject/internal/adapters/client/cryptocompare"
	"cryptoProject/internal/adapters/storage/postgres"
	"cryptoProject/internal/cases"
	"cryptoProject/internal/ports/http/public"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"log"
)

type App struct {
	client  *cryptocompare.CryptoCompareClient
	storage *postgres.PgStorage
	service *cases.Service
	server  *public.Server
	cron    *cron.Cron
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {
	if err := a.loadConfig(); err != nil {
		return err
	}

	ctx = context.TODO()

	defaultTitles := viper.GetStringSlice("client.default_titles")
	connString := viper.GetString("postgresql.conn_string")

	client, err := cryptocompare.NewCryptoCompareClient(defaultTitles)
	if err != nil {
		return err
	}
	a.client = client

	storage, err := postgres.NewPgStorage(ctx, connString)
	if err != nil {
		return err
	}
	a.storage = storage

	service, err := cases.NewService(storage, client)
	if err != nil {
		return err
	}
	a.service = service

	server, err := public.NewServer(service)
	if err != nil {
		return err
	}
	a.server = server
	server.Run()

	a.cron = cron.New()
	_, err = a.cron.AddFunc("@every 5m", func() {
		if err := a.service.UpdateCoinRates(ctx); err != nil {
			log.Printf("Failed to update coins rates: %v", err)
		}
	})
	if err != nil {
		return err
	}
	a.cron.Start()
	<-ctx.Done()
	a.cron.Stop()

	return nil
}

func (a *App) loadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// где вызывать NewCoin()??
