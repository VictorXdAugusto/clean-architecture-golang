package usecase

import (
	"go-crud/internal/application/dto"
	"go-crud/internal/application/repository"
)

type GetUser struct {
	UserRepository repository.User
}

func NewGetUser(
	userRepository repository.User,
) *GetUser {
	return &GetUser{
		UserRepository: userRepository,
	}
}

func (c *GetUser) Execute(id string) (*dto.GetUserOutput, error) {
	user, err := c.UserRepository.Get(id)

	if err != nil{
		return nil, err
	}

	return &dto.GetUserOutput{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
