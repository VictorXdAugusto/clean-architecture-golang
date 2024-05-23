package usecase

import (
	"go-crud/internal/application/dto"
	"go-crud/internal/application/repository"
)

type DeleteUser struct {
	UserRepository repository.User
}

func NewDeleteUser(
	userRepository repository.User,
) *DeleteUser {
	return &DeleteUser{
		UserRepository: userRepository,
	}
}

func (c *DeleteUser) Execute(id string) (*dto.DeleteUserOutput, error) {
	_ , err := c.UserRepository.Delete(id)
	
	if err != nil{
		return nil, err
	}

	return &dto.DeleteUserOutput{
		Message: "User successfully deleted",
	}, nil
}
