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
		Recipe_IDsTags:        map[uint][]string{1: {"sugar", "tea"}, 2: {"fish", "chips"}, 3: {"rice", "sushi"}, 4: {"vegaterian", "egg"}},
		Recipe_IDsDietlevel:   map[uint]uint{1: 1, 2: 2, 3: 2, 4: 1},
		All_Recipes_IDs:       []uint{1, 2, 3, 4},
		Recommendated_Recipes: []uint{}, // it is sorted by recommended points.
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

func TestDefinitelyRemoveDietLevels(t *testing.T) {
	recommendation := Construct()
	tests := []struct {
		output []uint
	}{
		{output: []uint{1, 4}},
	}
	for _, test := range tests {
		recommendation.DefinitelyRemoveDietLevels()

		sort.Slice(recommendation.All_Recipes_IDs, func(i, j int) bool {
			return recommendation.All_Recipes_IDs[i] < recommendation.All_Recipes_IDs[j]
		})

		if !reflect.DeepEqual(test.output, recommendation.All_Recipes_IDs) {
			t.Errorf("Result is: %v . Expected: %v", recommendation.All_Recipes_IDs, test.output)
		}
	}
	Destruct()
}
