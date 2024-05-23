package handler

import (
	"encoding/json"
	"go-crud/internal/application/dto"
	"go-crud/internal/application/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type UpdateUserHandler struct {
	UpdateUserInterface usecase.UpdateUserInteface
}

func (h *UpdateUserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    input := dto.UpdateUserInput{}

    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

    vars := mux.Vars(r)

    id := vars["id"]

    response, err := h.UpdateUserInterface.Execute(id, input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }   

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}