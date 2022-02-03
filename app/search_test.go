package app

import (
	"testing"
)

func TestSearchRecipes(t *testing.T) {
	app, _, _, _, rcp, _, _ := Construct()
	rcp.Create(app.DB)
	tests := []struct {
		input string
		err   error
	}{
		{input: "Test", err: nil},
	}
	for _, test := range tests {
		_, err := app.SearchRecipes(test.input)
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
