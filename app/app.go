package app

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/anilkusc/flow-diet-backend/api"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// App method is the main struct for the application
type App struct {
	API api.Api
	DB  *sql.DB
}

func (app *App) Init() {
	app.API.Router = mux.NewRouter()
	app.API.InitRoutes()
	if os.Getenv("ENV") == "dev" {
		log.SetLevel(log.DebugLevel)
	} else if os.Getenv("ENV") == "stg" {
		log.SetLevel(log.InfoLevel)
	} else if os.Getenv("ENV") == "prod" {
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(log.TraceLevel)
	}
}

func (app *App) Start() {
	app.Init()
	log.Warn("Serving on: 8080")
	log.Fatal(http.ListenAndServe(":8080", app.API.Router))
}
