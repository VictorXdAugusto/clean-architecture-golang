package usecase

import "go-crud/internal/application/dto"

type DeleteUserInteface interface {
	Execute(id string) (*dto.DeleteUserOutput, error)
}
