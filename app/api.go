package app

import (
	"os"

	_ "github.com/anilkusc/flow-diet-backend/docs" // docs is generated by Swag CLI, you have to import it.

	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *App) InitRoutes() {
	app.Router.HandleFunc("/user/signup", app.SignupHandler).Methods("POST")
	app.Router.HandleFunc("/user/signin", app.SigninHandler).Methods("POST")
	app.Router.HandleFunc("/user/logout", app.Auth(app.LogoutHandler)).Methods("POST")
	app.Router.HandleFunc("/user/test", app.TestHandler).Methods("GET")
	if os.Getenv("ENV") != "prod" {
		app.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	}
}
