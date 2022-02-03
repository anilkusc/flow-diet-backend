package app

import (
	"encoding/json"

	"github.com/anilkusc/flow-diet-backend/pkg/recipe"
	"github.com/anilkusc/flow-diet-backend/pkg/search"
)

func (app *App) SearchRecipes(word string) (string, error) {
	var recipes []recipe.Recipe
	filteredRecipes := map[uint]string{}
	var search search.Search
	search.Word = word
	rcpe := recipe.Recipe{}
	recipes, err := rcpe.List(app.DB)
	if err != nil {
		return "", err
	}
	for _, rcp := range recipes {
		filteredRecipes[rcp.ID] = rcp.Name
	}

	resultRecipes, err := search.FindRecipesByName(filteredRecipes)
	if err != nil {
		return "", err
	}

	var getMatchedRecipes []recipe.Recipe

	for _, r := range recipes {
		for _, rr := range resultRecipes {
			if r.ID == rr {
				getMatchedRecipes = append(getMatchedRecipes, r)
				continue
			}
		}
	}

	getMatchedRecipesList, err := json.Marshal(getMatchedRecipes)
	if err != nil {
		return "", err
	}
	return string(getMatchedRecipesList), nil
}
