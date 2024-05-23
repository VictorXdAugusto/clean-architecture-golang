package di

import (
	"github.com/google/wire"
	repository2 "go-crud/internal/application/repository"
	"go-crud/internal/infrastructure/repository"
)

var provideUserRepositoryPostgres = wire.NewSet(
	repository.NewUserPostgress,
	wire.Bind(new(repository2.User), new(*repository.UserPostgress)),
)
