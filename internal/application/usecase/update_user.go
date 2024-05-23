package usecase

import (
	"go-crud/internal/application/dto"
	"go-crud/internal/application/repository"
	"go-crud/internal/domain/entity"
)

type UpdateUser struct {
	UserRepository repository.User
}

func NewUpdateUser(
	userRepository repository.User,
) *UpdateUser {
	return &UpdateUser{
		UserRepository: userRepository,
	}
}

func (c *UpdateUser) Execute(id string, input dto.UpdateUserInput) (*dto.UpdateUserOutput, error) {
	updateUser := entity.NewUser(input.Name, input.Email, input.Username, input.Password, input.Age)

	_, err := c.UserRepository.Update(id, *updateUser)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUserOutput{
		ID:       updateUser.ID,
		Name:     updateUser.Name,
		Email:    updateUser.Email,
		Username: updateUser.Username,
	}, nil
}
