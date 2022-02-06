package recommendation

import (
	"reflect"
	"sort"
	"testing"

	"github.com/joho/godotenv"
)

func Construct() Recommendation {
	godotenv.Load("../../.env")
	var recommendation = Recommendation{
		Users_Preferred_Meals: []string{"breakfast", "noon", "night"},
		Users_Prohibits:       []string{"sugar"},
		Users_Diet_Level:      1,
		Users_Dislikes:        []string{"onion", "tomato"},
		Users_Likes:           []string{"chicken", "fish"},
		Users_Cousines:        []string{"mediterrian", "american"},

		Recipe_IDsAppropriateMeals: map[uint][]string{1: {"breakfast", "night"}, 2: {"noon"}, 3: {"snack"}},
		Recipe_IDsTags:             map[uint][]string{1: {"sugar", "tea"}, 2: {"fish", "chips"}, 3: {"rice", "sushi"}, 4: {"vegaterian", "egg"}},
		Recipe_IDsCousines:         map[uint][]string{1: {"italian", "mediterrian"}, 2: {"american"}, 3: {"mexico"}, 4: {"mediterrian", "syria"}},
		Recipe_IDsDietlevel:        map[uint]uint{1: 1, 2: 2, 3: 2, 4: 1},
		All_Recipes_IDs:            []uint{1, 2, 3, 4},
		Recipe_IDsPoints:           map[uint]uint{},

		Meal_Factor:         2,
		Like_Factor:         3,
		Dislike_Factor:      2,
		Cousine_Factor:      2,
		Recommended_Recipes: []uint{}, // it is sorted by recommended points.
	}
	return recommendation
}
func Destruct() {
	//db.Exec("DROP TABLE recipes")
}

func TestDefinitelyRemoveProhibits(t *testing.T) {
	recommendation := Construct()
	tests := []struct {
		output []uint
	}{
		{output: []uint{2, 3, 4}},
	}
	for _, test := range tests {
		recommendation.DefinitelyRemoveProhibits()

		sort.Slice(recommendation.All_Recipes_IDs, func(i, j int) bool {
			return recommendation.All_Recipes_IDs[i] < recommendation.All_Recipes_IDs[j]
		})

		if !reflect.DeepEqual(test.output, recommendation.All_Recipes_IDs) {
			t.Errorf("Result is: %v . Expected: %v", recommendation.All_Recipes_IDs, test.output)
		}
	}
	Destruct()
}
func TestRemoveFromAllRecipesIDs(t *testing.T) {
	recommendation := Construct()
	tests := []struct {
		input  uint
		output []uint
	}{
		{input: 1, output: []uint{2, 3, 4}},
	}
	for _, test := range tests {
		recommendation.RemoveFromAllRecipesIDs(test.input)

		sort.Slice(recommendation.All_Recipes_IDs, func(i, j int) bool {
			return recommendation.All_Recipes_IDs[i] < recommendation.All_Recipes_IDs[j]
		})

		if !reflect.DeepEqual(test.output, recommendation.All_Recipes_IDs) {
			t.Errorf("Result is: %v . Expected: %v", recommendation.All_Recipes_IDs, test.output)
		}
	}
	Destruct()
}

func TestDefinitelyRemoveHigherDietLevels(t *testing.T) {
	recommendation := Construct()
	tests := []struct {
		output []uint
	}{
		{output: []uint{1, 4}},
	}
	for _, test := range tests {
		recommendation.DefinitelyRemoveHigherDietLevels()

		sort.Slice(recommendation.All_Recipes_IDs, func(i, j int) bool {
			return recommendation.All_Recipes_IDs[i] < recommendation.All_Recipes_IDs[j]
		})

		if !reflect.DeepEqual(test.output, recommendation.All_Recipes_IDs) {
			t.Errorf("Result is: %v . Expected: %v", recommendation.All_Recipes_IDs, test.output)
		}
	}
	Destruct()
}
func TestPointByMeals(t *testing.T) {
	recommendation := Construct()
	tests := []struct {
		output map[uint]uint
	}{
		{output: map[uint]uint{1: 14, 2: 5, 3: 3}},
	}
	for _, test := range tests {
		recommendation.PointByMeals()

		if !reflect.DeepEqual(test.output, recommendation.Recipe_IDsPoints) {
			t.Errorf("Result is: %v . Expected: %v", recommendation.Recipe_IDsPoints, test.output)
		}
	}
	Destruct()
}
func TestPointByLikes(t *testing.T) {
	recommendation := Construct()
	tests := []struct {
		output map[uint]uint
	}{
		{output: map[uint]uint{1: 4, 2: 8, 3: 4, 4: 4}},
	}
	for _, test := range tests {
		recommendation.PointByLikes()

		if !reflect.DeepEqual(test.output, recommendation.Recipe_IDsPoints) {
			t.Errorf("Result is: %v . Expected: %v", recommendation.Recipe_IDsPoints, test.output)
		}
	}
	Destruct()
}
func TestPointByDislikes(t *testing.T) {
	recommendation := Construct()
	tests := []struct {
		output map[uint]uint
	}{
		{output: map[uint]uint{1: 4, 2: 3, 3: 4, 4: 4}},
	}
	for _, test := range tests {
		recommendation.PointByDislikes()

		if !reflect.DeepEqual(test.output, recommendation.Recipe_IDsPoints) {
			t.Errorf("Result is: %v . Expected: %v", recommendation.Recipe_IDsPoints, test.output)
		}
	}
	Destruct()
}
func TestReverseSortRecipeIdsByPoint(t *testing.T) {
	recommendation := Construct()
	recommendation.Recipe_IDsPoints = map[uint]uint{1: 10, 2: 20, 3: 15, 4: 2, 5: 12, 6: 3}
	tests := []struct {
		output []uint
	}{
		{output: []uint{2, 3, 5, 1, 6, 4}},
	}
	for _, test := range tests {
		res := recommendation.ReverseSortRecipeIdsByPoint()

		if !reflect.DeepEqual(test.output, res) {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}
	}
	Destruct()
}
func TestPointByCousine(t *testing.T) {
	recommendation := Construct()
	tests := []struct {
		output map[uint]uint
	}{
		{output: map[uint]uint{1: 7, 2: 4, 3: 2, 4: 5}},
	}
	for _, test := range tests {
		recommendation.PointByCousine()

		if !reflect.DeepEqual(test.output, recommendation.Recipe_IDsPoints) {
			t.Errorf("Result is: %v . Expected: %v", recommendation.Recipe_IDsPoints, test.output)
		}
	}
	Destruct()
}
