package app

import (
	"encoding/json"
	"testing"
)

func TestCreateMaterial(t *testing.T) {
	app, _, _, _, _, _, _, _, mtr := Construct()
	materialJson, _ := json.Marshal(mtr)
	tests := []struct {
		input string
		err   error
	}{
		{
			input: string(materialJson),
			err:   nil},
	}
	for _, test := range tests {
		err := app.CreateMaterial(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}

func TestGetMaterial(t *testing.T) {
	app, _, _, _, _, _, _, _, mtr := Construct()
	mtr.Create(app.DB)
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
		_, err := app.GetMaterial(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		} /*
			if output != test.output {
				t.Errorf("Result is: %v . Expected: %v", output, test.output)
			}*/
	}
	Destruct(app)
}

func TestDeleteMaterial(t *testing.T) {
	app, _, _, _, _, _, _, _, mtr := Construct()
	mtr.Create(app.DB)
	mtr.ID = 1
	materialJson, _ := json.Marshal(mtr)
	tests := []struct {
		input string
		err   error
	}{
		{
			input: string(materialJson),
			err:   nil},
	}
	for _, test := range tests {
		err := app.DeleteMaterial(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}

func TestUpdateMaterial(t *testing.T) {
	app, _, _, _, _, _, _, _, mtr := Construct()
	mtr.Create(app.DB)
	mtr.ID = 1
	mtr.Material_Name = "onion"
	materialJson, _ := json.Marshal(mtr)
	tests := []struct {
		input string
		err   error
	}{
		{
			input: string(materialJson),
			err:   nil},
	}
	for _, test := range tests {
		err := app.UpdateMaterial(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}
