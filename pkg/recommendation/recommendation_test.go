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
		Prohibits: []string{"sugar"},
	}
	return recommendation
}
func Destruct() {
	//db.Exec("DROP TABLE recipes")
}

func TestRemoveProhibits(t *testing.T) {
	recommendation := Construct()
	tests := []struct {
		input  map[uint]string
		output []uint
		err    error
	}{
		{input: map[uint]string{1: "sugar", 2: "banana", 3: "chicken", 4: "fish"}, output: []uint{2, 3, 4}, err: nil},
	}
	for _, test := range tests {
		res, err := recommendation.RemoveProhibits(test.input)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		sort.Slice(res, func(i, j int) bool {
			return test.output[i] < test.output[j]
		})
		if !reflect.DeepEqual(test.output, res) {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}
	}
	Destruct()
}
