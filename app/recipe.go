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

func (app *App) CreateRecipe(recipeJson string) error {
	var recipe recipe.Recipe
	var err error
	err = json.Unmarshal([]byte(recipeJson), &recipe)
	if err != nil {
		return err
	}

	err = recipe.Create(app.DB)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) GetRecipe(recipeJson string) (string, error) {

	var recipe recipe.Recipe
	var err error
	err = json.Unmarshal([]byte(recipeJson), &recipe)
	if err != nil {
		return "", err
	}

	recipe.Read(app.DB)
	if err != nil {
		return "", err
	}
	recipeStr, err := json.Marshal(recipe)
	if err != nil {
		return "", err
	}
	return string(recipeStr), nil
}
