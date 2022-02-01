package calendar

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, calendar := Construct()

	tests := []struct {
		input Calendar
		err   error
	}{
		{input: calendar, err: nil},
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
	db, calendar := Construct()
	calendar.Create(db)
	tests := []struct {
		input  Calendar
		output Calendar
		err    error
	}{
		{input: calendar, output: calendar, err: nil},
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
	db, calendar := Construct()
	calendar.Create(db)

	tests := []struct {
		input Calendar
		err   error
	}{
		{input: calendar, err: nil},
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
	db, calendar := Construct()
	calendar.Create(db)
	tests := []struct {
		input Calendar
		err   error
	}{
		{input: Calendar{
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
	db, calendar := Construct()
	calendar.Create(db)
	tests := []struct {
		input Calendar
		err   error
	}{
		{input: Calendar{
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
	db, calendar := Construct()
	calendar.Create(db)

	tests := []struct {
		input  Calendar
		output []Calendar
		err    error
	}{
		{

			input:  Calendar{Model: gorm.Model{ID: 1}},
			output: []Calendar{calendar}, err: nil},
	}
	for _, test := range tests {

		res, err := calendar.List(db)
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
