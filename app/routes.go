package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	user "github.com/anilkusc/flow-diet-backend/pkg/user"
	log "github.com/sirupsen/logrus"
)

// SignupHandler godoc
// @Summary Signup User
// @Description Create a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.User true "Create New User"
// @Success 200
// @Router /user/signup [post]
func (app *App) SignupHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}
	err = app.Signup(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	log.Info("user has been created: ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// SigninHandler godoc
// @Summary Signin User
// @Description Sign in with specified user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.User true "Sign In"
// @Success 200
// @Router /user/signin [post]
func (app *App) SigninHandler(w http.ResponseWriter, r *http.Request) {
	//this sleep is for preventing brute force
	time.Sleep(1 * time.Second)
	var user user.User
	var isauth bool
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}
	user, isauth, err = app.Signin(string(body))
	if err != nil {
		log.Error("cannot signin : ", err)
		http.Error(w, "cannot signin", http.StatusInternalServerError)
		return
	}
	if !isauth {
		log.Info("invalid credentials: ", string(body))
		http.Error(w, "invalid credentials", http.StatusForbidden)
		return
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Error("cannot marshall calendars json: ", err)
		http.Error(w, "cannot signin", http.StatusInternalServerError)
		return
	}
	session, err := app.SessionStore.Get(r, "session")
	if err != nil {
		log.Error("cannot get session store : ", err)
		http.Error(w, "cannot get session store", http.StatusInternalServerError)
		return
	}
	user.Read(app.DB)
	log.Info("updating session")
	session.Values["authenticated"] = "true"
	session.Values["role"] = user.Role
	//session.Values["id"] = strconv.FormatUint(uint64(user.ID), 10)
	session.Values["id"] = user.ID
	session.Save(r, w)
	log.Info("session updated")
	log.Info("user has been logged in: ", string(userJson))
	http.Error(w, string(userJson), http.StatusOK)
	return
}

// LogoutHandler godoc
// @Summary Logout User
// @Description Logout for the user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200
// @Router /user/logout [post]
func (app *App) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.SessionStore.Get(r, "session")
	if err != nil {
		log.Error("cannot get session store : ", err)
		http.Error(w, "cannot get session store", http.StatusInternalServerError)
		return
	}
	session.Values["authenticated"] = "false"
	session.Save(r, w)
	http.Error(w, "OK", http.StatusOK)
	return
}

// GetRecipesHandler godoc
// @Summary Get recipes of user weekly
// @Description Get recipes of the user weekly
// @Tags calendar
// @Accept  json
// @Produce  json
// @Success 200
// @Router /calendar/recipes [get]
func (app *App) GetCalendarRecipesHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.SessionStore.Get(r, "session")
	if err != nil {
		log.Error("cannot get session store : ", err)
		http.Error(w, "cannot get session store", http.StatusInternalServerError)
		return
	}
	calendars, err := app.GetMyCalendar(session.Values["id"].(uint))
	if err != nil {
		log.Error(err)
		http.Error(w, "cannot get recipes", http.StatusInternalServerError)
		return
	}
	calendarsJson, err := json.Marshal(calendars)
	if err != nil {
		log.Error("cannot marshall calendars json : ", err)
		http.Error(w, "Calendar Error", http.StatusInternalServerError)
		return
	}
	log.Info("user got his calendar: ", string(calendarsJson))
	http.Error(w, string(calendarsJson), http.StatusOK)
	return
}

// CreateCalendarRecipeHandler godoc
// @Summary Create Recipe In User Calendar
// @Description User creates a recipe in the calendar
// @Tags calendar
// @Accept  json
// @Produce  json
// @Param calendar body calendar.Calendar true "Create Recipe In Calendar"
// @Success 200
// @Router /calendar/recipes/create [post]
func (app *App) CreateCalendarRecipeHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong calendar object", http.StatusBadRequest)
		return
	}
	session, err := app.SessionStore.Get(r, "session")
	if err != nil {
		log.Error("cannot get session store : ", err)
		http.Error(w, "cannot get session store", http.StatusInternalServerError)
		return
	}
	err = app.CreateCalendar(string(body), session.Values["id"].(uint))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("created recipe in the calendar: ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// UpdateCalendarRecipeHandler godoc
// @Summary Update Recipe In User Calendar
// @Description Update Recipe In User Calendar
// @Tags calendar
// @Accept  json
// @Produce  json
// @Param calendar body calendar.Calendar true "Update Recipe In Calendar"
// @Success 200
// @Router /calendar/recipes/update [post]
func (app *App) UpdateCalendarRecipeHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong calendar object", http.StatusBadRequest)
		return
	}
	session, err := app.SessionStore.Get(r, "session")
	if err != nil {
		log.Error("cannot get session store : ", err)
		http.Error(w, "cannot get session store", http.StatusInternalServerError)
		return
	}
	err = app.UpdateCalendar(string(body), session.Values["id"].(uint))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("updated recipe in the calendar: ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// DeleteCalendarRecipeHandler godoc
// @Summary Delete Recipe In User Calendar
// @Description Delete Recipe In User Calendar
// @Tags calendar
// @Accept  json
// @Produce  json
// @Param calendar body calendar.Calendar true "Delete Recipe In Calendar. Please Use thisfor send request: {'ID':1}"
// @Success 200
// @Router /calendar/recipes/delete [post]
func (app *App) DeleteCalendarRecipeHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong calendar object", http.StatusBadRequest)
		return
	}
	session, err := app.SessionStore.Get(r, "session")
	if err != nil {
		log.Error("cannot get session store : ", err)
		http.Error(w, "cannot get session store", http.StatusInternalServerError)
		return
	}
	err = app.DeleteCalendar(string(body), session.Values["id"].(uint))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("deleted recipe in the calendar: ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// GetAllRecipesHandler godoc
// @Summary List all recipes
// @Description List All Recipes
// @Tags recipe
// @Accept  json
// @Produce  json
// @Success 200
// @Router /recipes/all [get]
func (app *App) GetAllRecipesHandler(w http.ResponseWriter, r *http.Request) {

	recipes, err := app.ListRecipes()
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("recipes are listed: ", recipes)
	http.Error(w, recipes, http.StatusOK)
	return
}

func (app *App) TestHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Hello", http.StatusOK)
	return
}
