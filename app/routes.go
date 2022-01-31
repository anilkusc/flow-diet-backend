package app

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// SignupHandler godoc
// @Summary Signup User
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body user.User true "Create New User"
// @Success 200
// @Router /user/signup [post]
func (app *App) SignupHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}
	err = app.Signup(string(body))
	if err != nil {
		log.Error("cannot signup : ", err)
		http.Error(w, "cannot signup", http.StatusInternalServerError)
		return
	}
	log.Info("user has been created: ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

func (app *App) ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// TODO add this process to the middleware
	id, ok := vars["id"]
	if !ok {
		log.Error("id is missing in URI")
	}
	io.WriteString(w, id)
	return
}
