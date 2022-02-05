package app

import (
	"encoding/json"

	"github.com/anilkusc/flow-diet-backend/pkg/recipe"
	"github.com/anilkusc/flow-diet-backend/pkg/recommendation"
	user "github.com/anilkusc/flow-diet-backend/pkg/user"
)

func (app *App) RecommendRecipes(userid uint) (string, error) {
	var recommendedRecipes []recipe.Recipe
	user := user.User{}
	user.ID = userid
	err := user.Read(app.DB)
	if err != nil {
		return "", err
	}
	recipe := recipe.Recipe{}
	recipes, err := recipe.List(app.DB)
	if err != nil {
		return "", err
	}

	recommendation := recommendation.Recommendation{
		Users_Preferred_Meals:      user.Preferred_Meals,
		Users_Prohibits:            user.Prohibits,
		Users_Diet_Level:           user.Diet_Level,
		Users_Dislikes:             user.Dislikes,
		Users_Likes:                user.Likes,
		Recipe_IDsPoints:           map[uint]uint{},
		Recipe_IDsAppropriateMeals: map[uint][]string{},
		Recipe_IDsTags:             map[uint][]string{},
		Recipe_IDsDietlevel:        map[uint]uint{},
		Meal_Factor:                2,
		Like_Factor:                3,
		Dislike_Factor:             2,
		Recommended_Recipes:        []uint{},
	}

	for _, rcp := range recipes {
		recommendation.Recipe_IDsAppropriateMeals[rcp.ID] = rcp.Appropriate_Meals
		recommendation.Recipe_IDsTags[rcp.ID] = rcp.Tags
		recommendation.Recipe_IDsDietlevel[rcp.ID] = rcp.Recipe_Diet_Level
		recommendation.All_Recipes_IDs = append(recommendation.All_Recipes_IDs, rcp.ID)
	}

	recommendation.MakeRecipeRecommendation()
	for _, recommendedRecipeID := range recommendation.Recommended_Recipes {
		for _, r := range recipes {
			if r.ID == recommendedRecipeID {
				recommendedRecipes = append(recommendedRecipes, r)
				break
			}
		}
	}
	recommendedRecipesListJson, err := json.Marshal(recommendedRecipes)
	if err != nil {
		return "", err
	}
	return string(recommendedRecipesListJson), nil

}
