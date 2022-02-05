package recommendation

type Recommendation struct {
	Users_Dislikes        []string
	Users_Likes           []string
	Users_Preferred_Meals []string
	Users_Prohibits       []string
	Users_Diet_Level      uint

	Recipe_IDsTags        map[uint][]string
	Recipe_IDsDietlevel   map[uint]uint
	All_Recipes_IDs       []uint
	Recommendated_Recipes []uint // it is sorted by recommended points.
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

func (r *Recommendation) DefinitelyRemoveDietLevels() {

	for recipeID, recipeDietlevel := range r.Recipe_IDsDietlevel {
		if recipeDietlevel > r.Users_Diet_Level {
			r.RemoveFromAllRecipesIDs(recipeID)
		}
	}
}

/*

func (r *Recommendation) PointByMeals(recipeIDsAppropriateMeals map[uint][]string) map[uint]int {
	recipeIDsPoint := map[uint]uint{}

	for recipeID, AppropriateMeals := range recipeIDsAppropriateMeals {
		point := 1
		for _, Users_Preferred_Meal := range r.Users_Preferred_Meals {
			for _, meal := range AppropriateMeals {
				if meal == Users_Preferred_Meal {
					point = point * 2
				}
			}

		}
		recipeIDsPoint[recipeID] = uint(point)
	}
	return recipeIDsPoint
}

func (r *Recommendation) PointByLikes(recipeIDsLikes map[uint][]string) map[uint]int {
	recipeIDsPoint := map[uint]uint{}

	for recipeID, likes := range recipeIDsLikes {
		point := 1
		for _, Users_Preferred_Meal := range r.Users_Preferred_Meals {
			for _, meal := range AppropriateMeals {
				if meal == Users_Preferred_Meal {
					point = point * 3
				}
			}

		}
		recipeIDsPoint[recipeID] = uint(point)
	}
	return recipeIDsPoint
}
*/
