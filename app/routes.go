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
	var usr user.User
	var isauth bool
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}
	usr, isauth, err = app.Signin(string(body))
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
	userJson, err := json.Marshal(usr)
	if err != nil {
		log.Error("cannot marshall user json: ", err)
		http.Error(w, "cannot signin", http.StatusInternalServerError)
		return
	}
	session, err := app.SessionStore.Get(r, "session")
	if err != nil {
		log.Error("cannot get session store : ", err)
		http.Error(w, "cannot get session store", http.StatusInternalServerError)
		return
	}

	log.Info("updating session")
	session.Values["authenticated"] = "true"
	session.Values["role"] = usr.Role
	session.Values["id"] = usr.ID
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
// @Summary Create Recipes In User Calendar
// @Description User creates a recipe in the calendar
// @Tags calendar
// @Accept  json
// @Produce  json
// @Param calendar body []calendar.Calendar true "Create Recipes In Calendar"
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
// @Router /calendar/recipes/update [patch]
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
// @Router /calendar/recipes/delete [delete]
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

// CreateRecipeHandler godoc
// @Summary Create a new recipe
// @Description Create A New Recipe
// @Tags recipe
// @Accept  json
// @Produce  json
// @Param recipe body recipe.Recipe true "Create New Recipe"
// @Success 200
// @Router /recipes/create [post]
func (app *App) CreateRecipeHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong calendar object", http.StatusBadRequest)
		return
	}

	err = app.CreateRecipe(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("recipe is create: ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// GetRecipeHandler godoc
// @Summary Get a recipe
// @Description  Get a recipe
// @Tags recipe
// @Accept  json
// @Produce  json
// @Param recipe body recipe.Recipe true "Get Recipe"
// @Success 200
// @Router /recipes/get [post]
func (app *App) GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong calendar object", http.StatusBadRequest)
		return
	}
	recipe, err := app.GetRecipe(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("recipe is readed: ", recipe)
	http.Error(w, recipe, http.StatusOK)
	return
}

// UpdateRecipeHandler godoc
// @Summary Update Recipe
// @Description Update Recipe
// @Tags recipe
// @Accept  json
// @Produce  json
// @Param recipe body recipe.Recipe true "Update a Recipe"
// @Success 200
// @Router /recipes/update [patch]
func (app *App) UpdateRecipeHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong calendar object", http.StatusBadRequest)
		return
	}

	err = app.UpdateRecipe(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("updated recipe : ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// DeleteRecipeHandler godoc
// @Summary Delete Recipe
// @Description Delete Recipe
// @Tags recipe
// @Accept  json
// @Produce  json
// @Param recipe body recipe.Recipe true "Delete a Recipe"
// @Success 200
// @Router /recipes/delete [delete]
func (app *App) DeleteRecipeHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong calendar object", http.StatusBadRequest)
		return
	}

	err = app.DeleteRecipe(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("deleted recipe : ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// GetAllShoppingsHandler godoc
// @Summary Get shopping lists
// @Description List All Shopping Lists
// @Tags shopping
// @Accept  json
// @Produce  json
// @Param shopping body shopping.Shopping true "Get shopping lists by date"
// @Success 200
// @Router /shopping/list [post]
func (app *App) GetAllShoppingsHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.SessionStore.Get(r, "session")
	if err != nil {
		log.Error("cannot get session store : ", err)
		http.Error(w, "cannot get session store", http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong calendar object", http.StatusBadRequest)
		return
	}

	shoppingList, err := app.ListShoppingsWithDateInterval(session.Values["id"].(uint), string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("shopping lists are listed: ", shoppingList)
	http.Error(w, shoppingList, http.StatusOK)
	return
}

// CreateShoppingHandler godoc
// @Summary Create a new shopping list
// @Description Create A New Shopping List
// @Tags shopping
// @Accept  json
// @Produce  json
// @Param shopping body shopping.Shopping true "Create New shopping List"
// @Success 200
// @Router /shopping/create [post]
func (app *App) CreateShoppingHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong shopping list object", http.StatusBadRequest)
		return
	}

	err = app.CreateShopping(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("shopping list is created: ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// GetShoppingHandler godoc
// @Summary Get a shopping list
// @Description  Get a shopping list
// @Tags shopping
// @Accept  json
// @Produce  json
// @Param shopping body shopping.Shopping true "Get Shopping List"
// @Success 200
// @Router /shopping/get [post]
func (app *App) GetShoppingHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong shopping list object", http.StatusBadRequest)
		return
	}
	recipe, err := app.GetShopping(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("shopping list is obtained: ", recipe)
	http.Error(w, recipe, http.StatusOK)
	return
}

// UpdateShoppingHandler godoc
// @Summary Update Shopping List
// @Description Update Shopping List
// @Tags shopping
// @Accept  json
// @Produce  json
// @Param shopping body shopping.Shopping true "Update Shopping List"
// @Success 200
// @Router /shopping/update [patch]
func (app *App) UpdateShoppingHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong calendar object", http.StatusBadRequest)
		return
	}

	err = app.UpdateShopping(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("updated recipe : ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// DeleteShoppingHandler godoc
// @Summary Delete Shopping List
// @Description Delete Shopping List
// @Tags shopping
// @Accept  json
// @Produce  json
// @Param shopping body shopping.Shopping true "Delete Shopping List"
// @Success 200
// @Router /shopping/delete [delete]
func (app *App) DeleteShoppingHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong calendar object", http.StatusBadRequest)
		return
	}

	err = app.DeleteShopping(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("deleted recipe : ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// SearchRecipesHandler godoc
// @Summary Search Recipes
// @Description Search Recipes by Title
// @Tags search
// @Accept  json
// @Produce  json
// @Param search body string true "Please write search word directly"
// @Success 200
// @Router /search/recipes [post]
func (app *App) SearchRecipesHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong search object", http.StatusBadRequest)
		return
	}

	recipes, err := app.SearchRecipes(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("searched for "+string(body)+" and got this recipes : ", string(body))
	http.Error(w, recipes, http.StatusOK)
	return
}

// GetRecommendationsHandler godoc
// @Summary Get Recommendations
// @Description Get Recommendations
// @Tags recommendation
// @Accept  json
// @Produce  json
// @Param recommendation body string true "dates(epoch format) : {'start_date':1643914403 ,'end_date':1644173603 }"
// @Success 200
// @Router /recommendation/getrecipes [post]
func (app *App) GetRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.SessionStore.Get(r, "session")
	if err != nil {
		log.Error("cannot get session store : ", err)
		http.Error(w, "cannot get session store", http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong search object", http.StatusBadRequest)
		return
	}

	recipes, err := app.RecommendRecipes(session.Values["id"].(uint), string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("shopping lists are listed: ", recipes)
	http.Error(w, recipes, http.StatusOK)
	return
}

// GetMaterialHandler godoc
// @Summary Get a material
// @Description  Get a material
// @Tags material
// @Accept  json
// @Produce  json
// @Param material body material.Material true "Get Material"
// @Success 200
// @Router /materials/get [post]
func (app *App) GetMaterialHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong material object", http.StatusBadRequest)
		return
	}
	material, err := app.GetMaterial(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("material is readed: ", material)
	http.Error(w, material, http.StatusOK)
	return
}

// CreateMaterialHandler godoc
// @Summary Create a new material
// @Description Create A New Material
// @Tags material
// @Accept  json
// @Produce  json
// @Param recipe body material.Material true "Create New Material"
// @Success 200
// @Router /materials/create [post]
func (app *App) CreateMaterialHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong material object", http.StatusBadRequest)
		return
	}

	err = app.CreateMaterial(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("material is created: ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// UpdateMaterialHandler godoc
// @Summary Update Material
// @Description Update Material
// @Tags material
// @Accept  json
// @Produce  json
// @Param material body material.Material true "Update a Material"
// @Success 200
// @Router /materials/update [patch]
func (app *App) UpdateMaterialHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong material object", http.StatusBadRequest)
		return
	}

	err = app.UpdateMaterial(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("updated material : ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

// DeleteMaterialHandler godoc
// @Summary Delete Material
// @Description Delete Material
// @Tags material
// @Accept  json
// @Produce  json
// @Param material body material.Material true "Delete a Material"
// @Success 200
// @Router /materials/delete [delete]
func (app *App) DeleteMaterialHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("cannot read body: ", err)
		http.Error(w, "wrong calendar object", http.StatusBadRequest)
		return
	}

	err = app.DeleteMaterial(string(body))
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	log.Info("deleted material : ", string(body))
	http.Error(w, "OK", http.StatusOK)
	return
}

func (app *App) TestHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Hello", http.StatusOK)
	return
}
