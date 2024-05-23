package usecase

import "go-crud/internal/application/dto"

type CreateUserInteface interface {
	Execute(input dto.CreateUserInput) (dto.CreateUserOutput, error)
}
