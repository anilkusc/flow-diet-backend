package api

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func IdControl(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, ok := vars["id"]
		if !ok {
			log.Info("id is missing in URI")
			http.Error(w, "id is missing in URI", http.StatusBadRequest)
			return
		}
		next(w, r)
	}
}
