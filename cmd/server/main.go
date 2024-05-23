package main

import (
	"go-crud/di"
)

func main() {
	app, cleanup, err := di.InitializeApi()
	if err != nil {
		panic(err)
	}

	err = app.Start()
	if err != nil {
		panic(err)
	}

	cleanup()
}
