package app

import (
	"net/http"

	user "github.com/anilkusc/flow-diet-backend/pkg/user"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (app *App) IdControl(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := app.SessionStore.Get(r, "session")
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			log.Info("id is missing in Path")
			http.Error(w, "id is missing in Path", http.StatusBadRequest)
			return
		}
		if id != session.Values["id"] {
			log.Info("user sent wrong id")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
func (app *App) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := app.SessionStore.Get(r, "session")
		if session.Values["authenticated"] != "true" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return

		} else {
			next(w, r)
		}
	}
}
func (app *App) Authz(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := app.SessionStore.Get(r, "session")
		user := user.User{}
		user.ID = session.Values["id"].(uint)
		user.Read(app.DB)

		if user.Role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return

		} else {
			next(w, r)
		}
	}
}
