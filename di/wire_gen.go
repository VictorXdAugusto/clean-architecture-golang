// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"go-crud/internal/application/usecase"
	"go-crud/internal/infrastructure/api"
	"go-crud/internal/infrastructure/api/handler"
	"go-crud/internal/infrastructure/repository"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func InitializeApi() (*api.Application, func(), error) {
	helloHandler := &handler.HelloHandler{}
	db, cleanup, err := providePostgresClient()
	if err != nil {
		return nil, nil, err
	}
	userPostgress := repository.NewUserPostgress(db)
	createUser := usecase.NewCreateUser(userPostgress)
	createUserHandler := &handler.CreateUserHandler{
		CreateUserInterface: createUser,
	}
	getUser := usecase.NewGetUser(userPostgress)
	getUserHandler := &handler.GetUserHandler{
		GetUserInterface: getUser,
	}
	deleteUser := usecase.NewDeleteUser(userPostgress)
	deleteUserHandler := &handler.DeleteUserHandler{
		DeleteUserInterface: deleteUser,
	}
	updateUser := usecase.NewUpdateUser(userPostgress)
	updateUserHandler := &handler.UpdateUserHandler{
		UpdateUserInterface: updateUser,
	}
	application := &api.Application{
		Hello:      helloHandler,
		CreateUser: createUserHandler,
		GetUser:    getUserHandler,
		DeleteUser: deleteUserHandler,
		UpdateUser: updateUserHandler,
	}
	return application, func() {
		cleanup()
	}, nil
}

// wire.go:

var wireApiSet = wire.NewSet(
	providePostgresClient,

	provideUserRepositoryPostgres,

	provideCreateUserUseCaseSet,

	provideGetUseCaseSet,

	provideDeleteUseCaseSet,

	provideUpdateUseCaseSet,

	apiHandlersSet, wire.Struct(new(api.Application), "*"),
)
