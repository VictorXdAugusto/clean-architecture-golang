package di

import (
	"github.com/google/wire"
	"go-crud/internal/application/usecase"
)

var provideCreateUserUseCaseSet = wire.NewSet(
	usecase.NewCreateUser,
	wire.Bind(new(usecase.CreateUserInteface), new(*usecase.CreateUser)),
)

var provideGetUseCaseSet = wire.NewSet(
	usecase.NewGetUser,
	wire.Bind(new(usecase.GetUserInteface), new(*usecase.GetUser)),
)

var provideDeleteUseCaseSet = wire.NewSet(
	usecase.NewDeleteUser,
	wire.Bind(new(usecase.DeleteUserInteface), new(*usecase.DeleteUser)),
)

var provideUpdateUseCaseSet = wire.NewSet(
	usecase.NewUpdateUser,
	wire.Bind(new(usecase.UpdateUserInteface), new(*usecase.UpdateUser)),
)