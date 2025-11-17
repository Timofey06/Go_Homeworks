package main

import (
	"hw3/internal/infrastructure/repository"
	"hw3/internal/infrastructure/server"
)

func main() {
	server.StartServer(":8080", repository.NewCSVRepository("testdb.csv"))
}
