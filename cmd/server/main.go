package main

import "go-crud/internal/infrastructure/api"

func main() {
	var app api.Application
	app.Start()
}