package main

import (
	"hw2/internal/client"
)

func main() {
	client.StartClient("http://localhost:8080", "Hellow, world!")
}
