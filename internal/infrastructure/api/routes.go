package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *Application) SetupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", a.Hello.Hello).Methods("GET")
	r.HandleFunc("/user/create", a.CreateUser.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", a.GetUser.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}/delete", a.DeleteUser.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/{id}/update", a.UpdateUser.UpdateUser).Methods("PUT")
	http.Handle("/", r)
}
