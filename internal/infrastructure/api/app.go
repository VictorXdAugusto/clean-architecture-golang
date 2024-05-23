package api

import (
	"fmt"
	"go-crud/internal/infrastructure/api/handler"
	"net/http"
)

type Application struct {
	Hello      *handler.HelloHandler
	CreateUser *handler.CreateUserHandler
}

func (a *Application) Start() error {
	a.SetupRoutes()

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
		return err
	}

	return nil
}
