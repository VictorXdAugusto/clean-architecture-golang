package di

import (
	"github.com/google/wire"
	"go-crud/internal/application/usecase"
)

var provideCreateUserUseCaseSet = wire.NewSet(
	usecase.NewCreateUser,
	wire.Bind(new(usecase.CreateUserInteface), new(*usecase.CreateUser)),
)
