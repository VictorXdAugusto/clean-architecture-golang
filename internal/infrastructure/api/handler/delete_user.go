package handler

import (
	"encoding/json"
	"go-crud/internal/application/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type DeleteUserHandler struct {
	DeleteUserInterface usecase.DeleteUserInteface
}

func (h *DeleteUserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    vars := mux.Vars(r)

    id := vars["id"]

    response, err := h.DeleteUserInterface.Execute(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }   

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}