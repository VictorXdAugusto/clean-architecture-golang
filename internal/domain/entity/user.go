package entity

import (
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

func NewUser(name, email, username, password string, age int) *User {
	return &User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Age:       age,
		IsActive:  true,
	}
}
