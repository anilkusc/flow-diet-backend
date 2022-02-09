package app

import (
	"encoding/json"

	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient"
)

func (app *App) CreateIngredient(ingredientJson string) error {
	var ingredient ingredient.Ingredient
	var err error
	err = json.Unmarshal([]byte(ingredientJson), &ingredient)
	if err != nil {
		return err
	}

	err = ingredient.Create(app.DB)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) GetIngredient(ingredientJson string) (string, error) {

	var ingredient ingredient.Ingredient
	var err error
	err = json.Unmarshal([]byte(ingredientJson), &ingredient)
	if err != nil {
		return "", err
	}
	err = ingredient.Read(app.DB)
	if err != nil {
		return "", err
	}
	ingredientStr, err := json.Marshal(ingredient)
	if err != nil {
		return "", err
	}
	return string(ingredientStr), nil
}

func (app *App) UpdateIngredient(ingredientJson string) error {

	var ingredient ingredient.Ingredient
	err := json.Unmarshal([]byte(ingredientJson), &ingredient)
	if err != nil {
		return err
	}

	err = ingredient.Update(app.DB)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) DeleteIngredient(ingredientJson string) error {

	var ingredient ingredient.Ingredient
	err := json.Unmarshal([]byte(ingredientJson), &ingredient)
	if err != nil {
		return err
	}

	err = ingredient.Delete(app.DB)
	if err != nil {
		return err
	}

	return nil
}
