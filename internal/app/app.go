package app

import (
	"context"
	"cryptoProject/internal/adapters/client/cryptocompare"
	"cryptoProject/internal/adapters/storage/postgres"
	"cryptoProject/internal/cases"
	"cryptoProject/internal/ports/http/public"
	cron2 "github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Run() {
	if err := loadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()
	defaultTitles := viper.GetStringSlice("client.default_titles")
	connString := viper.GetString("postgresql.conn_string")

	storage, err := postgres.NewPgStorage(ctx, connString)
	if err != nil {
		log.Fatalf("Failed to create storage: %v", err)
	}

	client, err := cryptocompare.NewCryptoCompareClient(defaultTitles)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	service, err := cases.NewService(storage, client)
	if err != nil {
		log.Fatalf("Failed to create service: %v", err)
	}

	go runCron(service)

	server, err := public.NewServer(service)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	server.Run()
	log.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func runCron(service *cases.Service) {
	cron := cron2.New()
	ctx := context.Background()
	updateFunc := func() {
		err := service.UpdateCoinRates(ctx)
		if err != nil {
			log.Printf("Failed to update coin rates: %v", err)
		}
	}
	cron.AddFunc("@every 30s", updateFunc)
	cron.Run()
}

func loadConfig() error {
	viper.SetConfigFile("./deployment/configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
