package entity

import (
	"go-crud/internal/application/dto"
	"time"

	"github.com/google/uuid"
)

// User é uma entidade que representa um usuário
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Age       int       `json:"age"`
	IsActive  bool      `json:"is_active"`
}

func NewUser(input dto.CreateUserInput) *User {
	return &User{
		ID: uuid.New().String(),
		Name: input.Name,
		Email: input.Email,
		Username: input.Username,
		Password: input.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Age: input.Age,
		IsActive: true,
	}
}
