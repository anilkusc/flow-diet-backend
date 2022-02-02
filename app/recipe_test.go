package app

import (
	"encoding/json"
	"testing"
)

func TestListRecipes(t *testing.T) {
	app, _, _, rcp := Construct()
	rcp.Create(app.DB)
	tests := []struct {
		//output []recipe.Recipe
		err error
	}{
		{
			//output: []recipe.Recipe{rcp},
			err: nil},
	}
	for _, test := range tests {
		_, err := app.ListRecipes()
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		// Since resulst cannot be compared skip for now
		/*
			if !reflect.DeepEqual(output, test.output) {
				t.Errorf("Result is: %v . Expected: %v", output, test.output)
			}*/

	}
	Destruct(app)
}
func TestCreateRecipe(t *testing.T) {
	app, _, _, rcp := Construct()
	recipeJson, _ := json.Marshal(rcp)
	tests := []struct {
		input string
		err   error
	}{
		{
			input: string(recipeJson),
			err:   nil},
	}
	for _, test := range tests {
		err := app.CreateRecipe(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}
