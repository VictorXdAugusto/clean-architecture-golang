package handler

import (
	"encoding/json"
	"go-crud/internal/application/dto"
	"go-crud/internal/application/usecase"
	"net/http"
)

type CreateUserHandler struct {
	CreateUserInterface usecase.CreateUserInteface
}

func (h *CreateUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	input := dto.CreateUserInput{}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	output, err := h.CreateUserInterface.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
