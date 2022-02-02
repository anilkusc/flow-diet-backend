package app

import (
	"encoding/json"

	"github.com/anilkusc/flow-diet-backend/pkg/recipe"
)

func (app *App) ListRecipes() (string, error) {
	var recipes []recipe.Recipe
	var recipe recipe.Recipe
	var err error

	recipes, err = recipe.List(app.DB)
	if err != nil {
		return "", err
	}
	recipesList, err := json.Marshal(recipes)
	if err != nil {
		return "", err
	}
	return string(recipesList), nil
}
