package api

import (
	"io"
	"net/http"
)

func (api *Api) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "returnValue")
	return
}
