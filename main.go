package main

import (
	"context"
	_ "cryptoProject/docs"
	"cryptoProject/internal/app"
	"log"
)

// @title Crypto API
// @version 1.0
// @description This is a sample server Crypto server
// @host localhost:8080
// @BasePath /v1
func main() {
	ctx := context.Background()
	application := app.NewApp()
	if err := application.Run(ctx); err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}

	//слой app.го - собрать приложение, запустить все конструкторы (в опр порядке)
	//конфиг файл (лучше yaml)
	//Run - запуск конструкторов
	//	cron джоба для обновления данных каждые 5 минут
	//	тестить без бд не получится, нужен докер, чтобы поднимать бд в контейнере
	//	.докер файл + докер компоуз
	//	докер декстоп - прилога докера, дб запущена
}
