package api

import "net/http"

func (a *Application) SetupRoutes() {
	http.HandleFunc("/", a.Hello.Hello)
	http.HandleFunc("/user/create", a.CreateUser.CreateUser)
}
