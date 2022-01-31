package api

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// CreateUserHandler godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "Create User With Id"
// @Success 200
// @Router /user/{id}/create [get]
func (api *Api) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "returnValue")
	return
}
func (api *Api) ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// TODO add this process to the middleware
	id, ok := vars["id"]
	if !ok {
		log.Error("id is missing in URI")
	}
	io.WriteString(w, id)
	return
}
func (api *Api) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "returnValue")
	return
}
func (api *Api) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "returnValue")
	return
}
