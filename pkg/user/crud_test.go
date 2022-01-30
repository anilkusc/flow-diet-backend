package user

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, user := Construct()

	tests := []struct {
		input User
		err   error
	}{
		{input: user, err: nil},
	}
	for _, test := range tests {
		err := test.input.Create(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct()
}

func TestRead(t *testing.T) {
	db, user := Construct()
	user.Create(db)
	tests := []struct {
		input  uint
		output User
		err    error
	}{
		{input: 1, output: user, err: nil},
	}
	for _, test := range tests {
		testuser := User{Model: gorm.Model{ID: 1}}
		err := testuser.Read(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		testuser.CreatedAt, testuser.UpdatedAt, test.output.CreatedAt, test.output.UpdatedAt = time.Time{}, time.Time{}, time.Time{}, time.Time{}
		testuser.DeletedAt, test.output.DeletedAt = gorm.DeletedAt{Time: time.Time{}, Valid: false}, gorm.DeletedAt{Time: time.Time{}, Valid: false}
		if !reflect.DeepEqual(testuser, test.output) {
			t.Errorf("Result is: %v . Expected: %v", testuser, test.output)
		}
	}
	Destruct()
}
func TestUpdate(t *testing.T) {
	db, user := Construct()
	user.Create(db)

	tests := []struct {
		input User
		err   error
	}{
		{input: user, err: nil},
	}
	for _, test := range tests {
		err := test.input.Update(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}

	Destruct()
}
func TestDelete(t *testing.T) {
	db, user := Construct()
	user.Create(db)
	tests := []struct {
		input User
		err   error
	}{
		{input: User{
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
	Destruct()
}

func TestHardDelete(t *testing.T) {
	db, user := Construct()
	user.Create(db)
	tests := []struct {
		input User
		err   error
	}{
		{input: User{
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
	Destruct()
}

func TestList(t *testing.T) {
	db, user := Construct()
	user.Create(db)

	tests := []struct {
		input  User
		output []User
		err    error
	}{
		{

			input:  User{Model: gorm.Model{ID: 1}},
			output: []User{user}, err: nil},
	}
	for _, test := range tests {

		res, err := user.List(db)
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
	Destruct()
}
