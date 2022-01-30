package api

import "github.com/gorilla/mux"

type Api struct {
	Router *mux.Router
}

func (api *Api) InitRoutes() {
	api.Router.HandleFunc("/user/create", api.CreateUserHandler)
	api.Router.HandleFunc("/user/{id}/read", api.ReadUserHandler)
	api.Router.HandleFunc("/user/{id}/update", api.UpdateUserHandler)
	api.Router.HandleFunc("/user/{id}/delete", api.DeleteUserHandler)
}
