package di

import (
	"github.com/google/wire"
	"go-crud/internal/infrastructure/api/handler"
)

var apiHandlersSet = wire.NewSet(
	wire.Struct(new(handler.HelloHandler), "*"),
	wire.Struct(new(handler.CreateUserHandler), "*"),
	wire.Struct(new(handler.GetUserHandler), "*"),
	wire.Struct(new(handler.DeleteUserHandler), "*"),
	wire.Struct(new(handler.UpdateUserHandler), "*"),
)
