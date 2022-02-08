package ingredient

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, ingredient := Construct()

	tests := []struct {
		input Ingredient
		err   error
	}{
		{input: ingredient, err: nil},
	}
	for _, test := range tests {
		err := test.input.Create(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(db)
}

func TestRead(t *testing.T) {
	db, ingredient := Construct()
	ingredient.Create(db)
	ingredient.ID = 1
	tests := []struct {
		input  Ingredient
		output Ingredient
		err    error
	}{
		{input: ingredient, output: ingredient, err: nil},
	}
	for _, test := range tests {

		err := test.input.Read(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		test.input.CreatedAt, test.input.UpdatedAt, test.output.CreatedAt, test.output.UpdatedAt = time.Time{}, time.Time{}, time.Time{}, time.Time{}
		test.input.DeletedAt, test.output.DeletedAt = gorm.DeletedAt{Time: time.Time{}, Valid: false}, gorm.DeletedAt{Time: time.Time{}, Valid: false}
		if !reflect.DeepEqual(test.input, test.output) {
			t.Errorf("Result is: %v . Expected: %v", test.input, test.output)
		}
	}
	Destruct(db)
}
func TestUpdate(t *testing.T) {
	db, ingredient := Construct()
	ingredient.Create(db)

	tests := []struct {
		input Ingredient
		err   error
	}{
		{input: ingredient, err: nil},
	}
	for _, test := range tests {
		err := test.input.Update(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}

	Destruct(db)
}
func TestDelete(t *testing.T) {
	db, ingredient := Construct()
	ingredient.Create(db)
	tests := []struct {
		input Ingredient
		err   error
	}{
		{input: Ingredient{
			Model: gorm.Model{
				ID: 1,
			},
		}, err: nil},
	}
	for _, test := range tests {
		err := test.input.Delete(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(db)
}

func TestHardDelete(t *testing.T) {
	db, ingredient := Construct()
	ingredient.Create(db)
	tests := []struct {
		input Ingredient
		err   error
	}{
		{input: Ingredient{
			Model: gorm.Model{
				ID: 1,
			},
		}, err: nil},
	}
	for _, test := range tests {
		err := test.input.HardDelete(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(db)
}

func TestList(t *testing.T) {
	db, ingredient := Construct()
	ingredient.Create(db)

	tests := []struct {
		input  Ingredient
		output []Ingredient
		err    error
	}{
		{

			input:  Ingredient{Model: gorm.Model{ID: 1}},
			output: []Ingredient{ingredient}, err: nil},
	}
	for _, test := range tests {

		res, err := ingredient.List(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		for i := range res {
			res[i].CreatedAt, res[i].UpdatedAt, test.output[i].CreatedAt, test.output[i].UpdatedAt = time.Time{}, time.Time{}, time.Time{}, time.Time{}
			res[i].DeletedAt, test.output[i].DeletedAt = gorm.DeletedAt{Time: time.Time{}, Valid: false}, gorm.DeletedAt{Time: time.Time{}, Valid: false}

			if !reflect.DeepEqual(res[i], test.output[i]) {
				t.Errorf("Result is: %v . Expected: %v", res[i], test.output[i])
				t.Errorf("Result list is: %v . Expected list: %v", res, test.output)
			}
		}

	}
	Destruct(db)
}
func TestListWithLimit(t *testing.T) {
	db, ingredient := Construct()
	ingredient.Create(db)
	ingredient2 := Ingredient{}
	ingredient2 = ingredient
	ingredient2.ID = 2
	ingredient2.Create(db)
	ingredient3 := Ingredient{}
	ingredient3 = ingredient
	ingredient3.ID = 3
	ingredient3.Create(db)
	ingredient4 := Ingredient{}
	ingredient4 = ingredient
	ingredient4.ID = 4
	ingredient4.Create(db)

	tests := []struct {
		limit int
		err   error
	}{
		{limit: 2, err: nil},
	}

	for _, test := range tests {

		res, err := ingredient.ListWithLimit(db, test.limit)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if len(res) > test.limit {
			t.Errorf("Length is: %v . Expected: %v", len(res), test.limit)
		}

	}
	Destruct(db)
}
