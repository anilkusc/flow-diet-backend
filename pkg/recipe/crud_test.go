package recipe

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, recipe := Construct()

	tests := []struct {
		input Recipe
		err   error
	}{
		{input: recipe, err: nil},
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
	db, recipe := Construct()
	recipe.Create(db)
	recipe.ID = 1
	tests := []struct {
		input  Recipe
		output Recipe
		err    error
	}{
		{input: recipe, output: recipe, err: nil},
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
	db, recipe := Construct()
	recipe.Create(db)

	tests := []struct {
		input Recipe
		err   error
	}{
		{input: recipe, err: nil},
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
	db, recipe := Construct()
	recipe.Create(db)
	tests := []struct {
		input Recipe
		err   error
	}{
		{input: Recipe{
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
	db, recipe := Construct()
	recipe.Create(db)
	tests := []struct {
		input Recipe
		err   error
	}{
		{input: Recipe{
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
	db, recipe := Construct()
	recipe.Create(db)

	tests := []struct {
		input  Recipe
		output []Recipe
		err    error
	}{
		{

			input:  Recipe{Model: gorm.Model{ID: 1}},
			output: []Recipe{recipe}, err: nil},
	}
	for _, test := range tests {

		res, err := recipe.List(db)
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
	db, recipe := Construct()
	recipe.Create(db)
	recipe2 := Recipe{}
	recipe2 = recipe
	recipe2.ID = 2
	recipe2.Create(db)
	recipe3 := Recipe{}
	recipe3 = recipe
	recipe3.ID = 3
	recipe3.Create(db)
	recipe4 := Recipe{}
	recipe4 = recipe
	recipe4.ID = 4
	recipe4.Create(db)

	tests := []struct {
		limit int
		err   error
	}{
		{limit: 2, err: nil},
	}

	for _, test := range tests {

		res, err := recipe.ListWithLimit(db, test.limit)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if len(res) > test.limit {
			t.Errorf("Length is: %v . Expected: %v", len(res), test.limit)
		}

	}
	Destruct(db)
}
