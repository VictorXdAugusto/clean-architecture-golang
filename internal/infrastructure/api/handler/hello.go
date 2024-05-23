package handler

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
}

func (h *HelloHandler) Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello, World!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
