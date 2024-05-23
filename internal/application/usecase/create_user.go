package usecase

import (
	"go-crud/internal/application/dto"
	"go-crud/internal/application/repository"
	"go-crud/internal/domain/entity"
)

type CreateUser struct {
	UserRepository repository.User
}

func NewCreateUser(
	userRepository repository.User,
) *CreateUser {
	return &CreateUser{
		UserRepository: userRepository,
	}
}

func (c *CreateUser) Execute(input dto.CreateUserInput) (dto.CreateUserOutput, error) {
	user := entity.NewUser(input.Name, input.Email, input.Username, input.Password, input.Age)

	_, _ = c.UserRepository.Insert(*user)

	return dto.CreateUserOutput{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
