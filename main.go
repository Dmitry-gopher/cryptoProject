package main

import (
	_ "cryptoProject/docs"
	"cryptoProject/internal/app"
)

// @title Crypto API
// @version 1.0
// @description This is a sample server Crypto server
// @host localhost:8080
// @BasePath /v1
func main() {
	app.Run()
}
