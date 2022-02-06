package recommendation

type Recommendation struct {
	Users_Dislikes        []string
	Users_Likes           []string
	Users_Preferred_Meals []string
	Users_Prohibits       []string
	Users_Cousines        []string

	Users_Diet_Level           uint
	Users_Want                 string
	Recipe_IDsTags             map[uint][]string
	Recipe_IDsCousines         map[uint][]string
	Recipe_IDsAppropriateMeals map[uint][]string
	Recipe_IDsDietlevel        map[uint]uint
	Recipe_IDsPoints           map[uint]uint
	All_Recipes_IDs            []uint

	Meal_Factor    uint // this is factor importance weight of Meals while point the recipe. It will be multiple with point.
	Like_Factor    uint
	Dislike_Factor uint
	Cousine_Factor uint

	Recommended_Recipes []uint // it is sorted by recommended points.
}

func (r *Recommendation) MakeRecipeRecommendation() {
	r.DefinitelyRemoveProhibits()
	r.DefinitelyRemoveHigherDietLevels()
	r.PointByMeals()
	r.PointByLikes()
	r.PointByDislikes()
	r.PointByCousine()
	r.Recommended_Recipes = r.ReverseSortRecipeIdsByPoint()
}

func (r *Recommendation) DefinitelyRemoveProhibits() {

	for recipeID, tags := range r.Recipe_IDsTags {
		isTagProhbited := false
		for _, tag := range tags {
			for _, prohibit := range r.Users_Prohibits {
				if prohibit == tag {
					isTagProhbited = true
					break
				}
			}

		}
		if isTagProhbited {
			r.RemoveFromAllRecipesIDs(recipeID)
		}
	}
}

func (r *Recommendation) RemoveFromAllRecipesIDs(element uint) {
	var index int
	for i, recipeID := range r.All_Recipes_IDs {
		if recipeID == element {
			index = i
			break
		}
	}
	r.All_Recipes_IDs = append(r.All_Recipes_IDs[:index], r.All_Recipes_IDs[index+1:]...)
}

func (r *Recommendation) DefinitelyRemoveHigherDietLevels() {

	for recipeID, recipeDietlevel := range r.Recipe_IDsDietlevel {
		if recipeDietlevel > r.Users_Diet_Level {
			r.RemoveFromAllRecipesIDs(recipeID)
		}
	}
}

func (r *Recommendation) PointByMeals() {

	for recipeID, AppropriateMeals := range r.Recipe_IDsAppropriateMeals {
		for _, AppropriateMeal := range AppropriateMeals {
			for _, Users_Preferred_Meal := range r.Users_Preferred_Meals {
				//TODO add +1 only if it is zero .
				r.Recipe_IDsPoints[recipeID] = r.Recipe_IDsPoints[recipeID] + 1
				if Users_Preferred_Meal == AppropriateMeal {
					r.Recipe_IDsPoints[recipeID] = r.Recipe_IDsPoints[recipeID] * r.Meal_Factor
				}
			}
		}

	}
}
func (r *Recommendation) PointByLikes() {

	for recipeID, tags := range r.Recipe_IDsTags {
		for _, tag := range tags {
			for _, Users_Like := range r.Users_Likes {
				r.Recipe_IDsPoints[recipeID] = r.Recipe_IDsPoints[recipeID] + 1
				if Users_Like == tag {
					r.Recipe_IDsPoints[recipeID] = r.Recipe_IDsPoints[recipeID] * r.Like_Factor
				}
			}
		}

	}
}

func (r *Recommendation) PointByDislikes() {

	for recipeID, tags := range r.Recipe_IDsTags {
		for _, tag := range tags {
			for _, Users_Like := range r.Users_Likes {
				r.Recipe_IDsPoints[recipeID] = r.Recipe_IDsPoints[recipeID] + 1
				if Users_Like == tag {
					r.Recipe_IDsPoints[recipeID] = r.Recipe_IDsPoints[recipeID] / r.Dislike_Factor
				}
			}
		}

	}
}

func (r *Recommendation) PointByCousine() {

	for recipeID, cousines := range r.Recipe_IDsCousines {
		for _, cousine := range cousines {
			for _, Users_Cousine := range r.Users_Cousines {
				r.Recipe_IDsPoints[recipeID] = r.Recipe_IDsPoints[recipeID] + 1
				if Users_Cousine == cousine {
					r.Recipe_IDsPoints[recipeID] = r.Recipe_IDsPoints[recipeID] * r.Cousine_Factor
				}
			}
		}

	}
}
