package app

import (
	"database/sql"
	"flag"
	"net/http"
	"os"

	"github.com/anilkusc/flow-diet-backend/api"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	port  = flag.String("port", "8080", "Specifiy Port(default:8080)")
	https = flag.Bool("https", false, "Enable or disable https(default:false)")
	store = sessions.NewCookieStore([]byte(os.Getenv("STORE_KEY")))
)

// App method is the main struct for the application
type App struct {
	API api.Api
	DB  *sql.DB
}

func (app *App) Init() {
	if os.Getenv("ENV") == "" {
		godotenv.Load("../.env")
	}
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
	flag.Parse()
}

func (app *App) Start() {
	app.Init()

	if *https {
		log.Warn("Serving with TLS on: " + *port)
		log.Fatal(http.ListenAndServeTLS(":"+*port, "./certs/server.crt", "./certs/server.key", app.API.Router))
	} else {
		log.Warn("Serving on: " + *port)
		log.Fatal(http.ListenAndServe(":"+*port, app.API.Router))
	}

}
