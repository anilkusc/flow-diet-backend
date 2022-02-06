package search

import (
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

func Construct() Search {
	godotenv.Load("../../.env")
	var search = Search{
		Word: "",
	}
	//db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//db.AutoMigrate(&Recipe{})
	return search
}
func Destruct() {
	//db.Exec("DROP TABLE recipes")
}

func TestFindRecipesByName(t *testing.T) {
	search := Construct()
	search.Word = "some"
	tests := []struct {
		input  map[uint]string
		output []uint
		err    error
	}{
		{input: map[uint]string{1: "something", 2: "somehow", 3: "alright", 4: "locked"}, output: []uint{1, 2}, err: nil},
	}
	for _, test := range tests {
		res, err := search.FindRecipesByName(test.input)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		//TODO add sort here
		if !reflect.DeepEqual(test.output, res) {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}
	}
	Destruct()
}
