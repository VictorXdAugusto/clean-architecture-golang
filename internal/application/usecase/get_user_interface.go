package usecase

import "go-crud/internal/application/dto"

type GetUserInteface interface {
	Execute(id string) (*dto.GetUserOutput, error)
}
