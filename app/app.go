package app

import (
	"flag"
	"net/http"
	"os"

	"github.com/anilkusc/flow-diet-backend/pkg/calendar"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient/material"
	"github.com/anilkusc/flow-diet-backend/pkg/shopping"
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
)

// App method is the main struct for the application
type App struct {
	Router       *mux.Router
	DB           *gorm.DB
	SessionStore *sessions.CookieStore
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
		app.DB.AutoMigrate(&user.User{}, &calendar.Calendar{}, &recipe.Recipe{}, &shopping.Shopping{}, &material.Material{}, &ingredient.Ingredient{})
		log.Info("created")
	}
	// TODO: will be implemented on rolebased access
	/*u := user.User{Password: os.Getenv("ROOT_PASS")}
	pass, _ := u.HashPassword(u.Password)
	app.DB.Exec("INSERT INTO users (id, username, password,role) SELECT * FROM (SELECT '1','" + os.Getenv("ROOT_USER") + "','" + pass + "','root') AS tmp WHERE NOT EXISTS ( SELECT username FROM users WHERE username = '" + os.Getenv("ROOT_USER") + "');")
	*/
	log.Info("creating session store")
	app.SessionStore = sessions.NewCookieStore([]byte(os.Getenv("STORE_KEY")))
	log.Info("created")
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
