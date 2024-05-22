package handler

import "net/http"

type CreateUserHandler struct {
}

func (h *CreateUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
        // Se o método for POST, escreve "Create User" na resposta HTTP
		
    } else {
        // Se o método não for GET, retorna um erro 405 Method Not Allowed
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }

}
