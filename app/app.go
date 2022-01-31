package app

import (
	"flag"
	"net/http"
	"os"

	user "github.com/anilkusc/flow-diet-backend/pkg/user"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	port  = flag.String("port", "8080", "Specifiy Port(default:8080)")
	https = flag.Bool("https", false, "Enable or disable https(default:false)")
	store = sessions.NewCookieStore([]byte(os.Getenv("STORE_KEY")))
)

// App method is the main struct for the application
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (app *App) Init() {
	var err error
	if os.Getenv("ENV") == "" {
		godotenv.Load()
	}
	app.Router = mux.NewRouter()
	app.InitRoutes()
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
	if os.Getenv("DB_CONN") == "sqlite" {
		log.Info("connecting sqlite database")
		app.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		log.Info("connected")
		log.Info("creating tables")
		app.DB.AutoMigrate(&user.User{})
		log.Info("created")
	}
}

func (app *App) Start() {
	app.Init()

	if *https {
		log.Warn("Serving with TLS on: " + *port)
		log.Fatal(http.ListenAndServeTLS(":"+*port, "./certs/server.crt", "./certs/server.key", app.Router))
	} else {
		log.Warn("Serving on: " + *port)
		log.Fatal(http.ListenAndServe(":"+*port, app.Router))
	}

}
