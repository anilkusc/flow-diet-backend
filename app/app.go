package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/anilkusc/flow-diet-backend/api"
	"github.com/gorilla/mux"
)

// App method is the main struct for the application
type App struct {
	API api.Api
	DB  *sql.DB
}

func (app *App) Init() {
	app.API.Router = mux.NewRouter()
	app.API.InitRoutes()
}

func (app *App) Start() {
	app.Init()
	log.Println("Serving on: 8080")
	log.Fatal(http.ListenAndServe(":8080", app.API.Router))
}
