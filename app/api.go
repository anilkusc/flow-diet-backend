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
	app.Router.HandleFunc("/calendar/recipes", app.Auth(app.GetCalendarRecipesHandler)).Methods("GET")
	app.Router.HandleFunc("/calendar/recipes/create", app.Auth(app.CreateCalendarRecipeHandler)).Methods("POST")
	app.Router.HandleFunc("/calendar/recipes/update", app.Auth(app.UpdateCalendarRecipeHandler)).Methods("POST")
	app.Router.HandleFunc("/calendar/recipes/delete", app.Auth(app.DeleteCalendarRecipeHandler)).Methods("POST")
	app.Router.HandleFunc("/recipes/all", app.Auth(app.GetAllRecipesHandler)).Methods("GET")
	app.Router.HandleFunc("/recipes/get", app.Auth(app.GetRecipeHandler)).Methods("POST")
	app.Router.HandleFunc("/recipes/create", app.Auth(app.Authz(app.CreateRecipeHandler))).Methods("POST")
	app.Router.HandleFunc("/recipes/update", app.Auth(app.Authz(app.UpdateRecipeHandler))).Methods("POST")
	app.Router.HandleFunc("/recipes/delete", app.Auth(app.Authz(app.DeleteRecipeHandler))).Methods("POST")
	app.Router.HandleFunc("/shopping/all", app.Auth(app.GetAllShoppingsHandler)).Methods("GET")
	app.Router.HandleFunc("/shopping/get", app.Auth(app.GetShoppingHandler)).Methods("POST")
	app.Router.HandleFunc("/shopping/create", app.Auth(app.CreateShoppingHandler)).Methods("POST")
	app.Router.HandleFunc("/shopping/update", app.Auth(app.UpdateShoppingHandler)).Methods("POST")
	app.Router.HandleFunc("/shopping/delete", app.Auth(app.DeleteShoppingHandler)).Methods("POST")
	app.Router.HandleFunc("/search/recipes", app.Auth(app.SearchRecipesHandler)).Methods("POST")
	app.Router.HandleFunc("/recommendation/getrecipes", app.Auth(app.GetRecommendationsHandler)).Methods("GET")
	if os.Getenv("ENV") != "prod" {
		app.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	}
}
