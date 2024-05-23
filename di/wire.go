//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"go-crud/internal/infrastructure/api"
)

var wireApiSet = wire.NewSet(
	providePostgresClient,

	provideUserRepositoryPostgres,

	provideCreateUserUseCaseSet,

	apiHandlersSet,
	wire.Struct(new(api.Application), "*"),
)

func InitializeApi() (*api.Application, func(), error) {
	wire.Build(wireApiSet)
	return &api.Application{}, func() {}, nil
}
