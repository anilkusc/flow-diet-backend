package app

import (
	"encoding/json"

	"github.com/anilkusc/flow-diet-backend/pkg/recipe"
	"github.com/anilkusc/flow-diet-backend/pkg/recommendation"
	user "github.com/anilkusc/flow-diet-backend/pkg/user"
)

func (app *App) RecommendRecipes(userid uint, datesJson string) (string, error) {

	type TimeInterval struct {
		Start_date int64 `json:"start_date"`
		End_date   int64 `json:"end_date"`
	}
	var timeinterval TimeInterval
	err := json.Unmarshal([]byte(datesJson), &timeinterval)
	if err != nil {
		return "", err
	}

	user := user.User{}
	user.ID = userid
	err = user.Read(app.DB)
	if err != nil {
		return "", err
	}

	recommendation := recommendation.Recommendation{
		Users_Preferred_Meals:      user.Preferred_Meals,
		Users_Prohibits:            user.Prohibits,
		Users_Diet_Level:           user.Diet_Level,
		Users_Dislikes:             user.Dislikes,
		Users_Likes:                user.Likes,
		Users_Cousines:             user.Favorite_Cousines,
		Recipe_IDsPoints:           map[uint]uint{},
		Recipe_IDsAppropriateMeals: map[uint][]string{},
		Recipe_IDsCousines:         map[uint][]string{},
		Recipe_IDsTags:             map[uint][]string{},
		Recipe_IDsDietlevel:        map[uint]uint{},
		Meal_Factor:                2,
		Like_Factor:                3,
		Dislike_Factor:             2,

		Start_Date: timeinterval.Start_date,
		End_Date:   timeinterval.End_date,

		Needed_Recipe_Count: 0,
		Recommended_Recipes: []uint{},
	}

	var recommendedRecipes []recipe.Recipe
	recipe := recipe.Recipe{}
	recipes, err := recipe.List(app.DB)
	if err != nil {
		return "", err
	}

	for _, rcp := range recipes {
		recommendation.Recipe_IDsAppropriateMeals[rcp.ID] = rcp.Appropriate_Meals
		recommendation.Recipe_IDsTags[rcp.ID] = rcp.Tags
		recommendation.Recipe_IDsDietlevel[rcp.ID] = rcp.Recipe_Diet_Level
		recommendation.Recipe_IDsCousines[rcp.ID] = rcp.Cousines
		recommendation.All_Recipes_IDs = append(recommendation.All_Recipes_IDs, rcp.ID)
	}

	err = recommendation.MakeRecipeRecommendation()
	if err != nil {
		return "", err
	}

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
