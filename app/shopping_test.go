package app

import (
	"encoding/json"
	"testing"
)

func TestListShoppings(t *testing.T) {
	app, _, _, _, _, shp, _, _, _ := Construct()
	shp.Create(app.DB)
	tests := []struct {
		userid uint
		err    error
	}{
		{userid: 1, err: nil},
	}
	for _, test := range tests {
		_, err := app.ListShoppings(test.userid)
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
func TestCreateShopping(t *testing.T) {
	app, _, _, _, _, shp, _, _, _ := Construct()
	shoppingJson, _ := json.Marshal(shp)
	tests := []struct {
		input string
		err   error
	}{
		{
			input: string(shoppingJson),
			err:   nil},
	}
	for _, test := range tests {
		err := app.CreateShopping(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}

func TestGetShopping(t *testing.T) {
	app, _, _, _, _, shp, _, _, _ := Construct()
	shp.Create(app.DB)
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
		_, err := app.GetShopping(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		} /*
			if output != test.output {
				t.Errorf("Result is: %v . Expected: %v", output, test.output)
			}*/
	}
	Destruct(app)
}

func TestDeleteShopping(t *testing.T) {
	app, _, _, _, _, shp, _, _, _ := Construct()
	shp.Create(app.DB)
	shp.ID = 1
	shoppingJson, _ := json.Marshal(shp)
	tests := []struct {
		input string
		err   error
	}{
		{
			input: string(shoppingJson),
			err:   nil},
	}
	for _, test := range tests {
		err := app.DeleteShopping(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}

func TestUpdateShopping(t *testing.T) {
	app, _, _, _, _, shp, _, _, _ := Construct()
	shp.Create(app.DB)
	shp.ID = 1
	shp.Start_Date = 100
	shoppingJson, _ := json.Marshal(shp)
	tests := []struct {
		input string
		err   error
	}{
		{
			input: string(shoppingJson),
			err:   nil},
	}
	for _, test := range tests {
		err := app.UpdateShopping(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}

func TestListShoppingsWithDateInterval(t *testing.T) {
	app, _, _, _, _, shp, _, _, _ := Construct()
	shp.Start_Date = 1643937031
	shp.End_Date = 1644016231
	shp.Create(app.DB)
	tests := []struct {
		userid       uint
		shoppingJson string
		err          error
	}{

		{userid: 1, shoppingJson: `{"start_date":1643850631 ,"end_date":1644109831}`, err: nil},
	}
	for _, test := range tests {
		_, err := app.ListShoppingsWithDateInterval(test.userid, test.shoppingJson)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}
