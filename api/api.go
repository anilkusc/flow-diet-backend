package api

import "github.com/gorilla/mux"

type Api struct {
	Router *mux.Router
}

func (api *Api) InitRoutes() {
	api.Router.HandleFunc("/user/create", api.CreateUserHandler)
}
