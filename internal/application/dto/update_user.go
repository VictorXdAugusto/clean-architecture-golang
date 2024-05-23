package dto

type UpdateUserInput struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
}

type UpdateUserOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
}