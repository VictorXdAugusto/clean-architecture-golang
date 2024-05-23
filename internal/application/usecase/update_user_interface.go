package usecase

import "go-crud/internal/application/dto"

type UpdateUserInteface interface {
	Execute(id string, input dto.UpdateUserInput) (*dto.UpdateUserOutput, error)
}
