package app

import (
	"encoding/json"
	"testing"
)

func TestCreateIngredient(t *testing.T) {
	app, _, _, _, _, _, _, _, _, ingr := Construct()
	ingredientJson, _ := json.Marshal(ingr)
	tests := []struct {
		input string
		err   error
	}{
		{
			input: string(ingredientJson),
			err:   nil},
	}
	for _, test := range tests {
		err := app.CreateIngredient(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}

func TestGetIngredient(t *testing.T) {
	app, _, _, _, _, _, _, _, _, ingr := Construct()
	ingr.Create(app.DB)
	tests := []struct {
		input  string
		output string
		err    error
	}{
		{
			input: "{\"id\":1}",
			//output: "string",
			err: nil},
	}
	for _, test := range tests {
		_, err := app.GetIngredient(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		} /*
			if output != test.output {
				t.Errorf("Result is: %v . Expected: %v", output, test.output)
			}*/
	}
	Destruct(app)
}

func TestDeleteIngredient(t *testing.T) {
	app, _, _, _, _, _, _, _, _, ingr := Construct()
	ingr.Create(app.DB)
	ingr.ID = 1
	ingredientJson, _ := json.Marshal(ingr)
	tests := []struct {
		input string
		err   error
	}{
		{
			input: string(ingredientJson),
			err:   nil},
	}
	for _, test := range tests {
		err := app.DeleteIngredient(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}

func TestUpdateIngredient(t *testing.T) {
	app, _, _, _, _, _, _, _, _, ingr := Construct()
	ingr.Create(app.DB)
	ingr.ID = 1
	ingr.IsOptional = true
	ingredientJson, _ := json.Marshal(ingr)
	tests := []struct {
		input string
		err   error
	}{
		{
			input: string(ingredientJson),
			err:   nil},
	}
	for _, test := range tests {
		err := app.UpdateIngredient(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}
