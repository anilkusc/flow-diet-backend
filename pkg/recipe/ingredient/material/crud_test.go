package material

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, material := Construct()

	tests := []struct {
		input Material
		err   error
	}{
		{input: material, err: nil},
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
	db, material := Construct()
	material.Create(db)
	material.ID = 1
	tests := []struct {
		input  Material
		output Material
		err    error
	}{
		{input: material, output: material, err: nil},
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
	db, material := Construct()
	material.Create(db)

	tests := []struct {
		input Material
		err   error
	}{
		{input: material, err: nil},
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
	db, material := Construct()
	material.Create(db)
	tests := []struct {
		input Material
		err   error
	}{
		{input: Material{
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
	db, material := Construct()
	material.Create(db)
	tests := []struct {
		input Material
		err   error
	}{
		{input: Material{
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
	db, material := Construct()
	material.Create(db)

	tests := []struct {
		input  Material
		output []Material
		err    error
	}{
		{

			input:  Material{Model: gorm.Model{ID: 1}},
			output: []Material{material}, err: nil},
	}
	for _, test := range tests {

		res, err := material.List(db)
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
