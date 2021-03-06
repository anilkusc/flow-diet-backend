package user

import (
	"testing"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Construct() (*gorm.DB, User) {

	godotenv.Load("../../.env")
	var user = User{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Username:          "testuser1",
		Name:              "test user",
		Email:             "testmail@test.com",
		Phone:             "+905355353535",
		Password:          "testpass",
		Weight:            70,
		Height:            173,
		Age:               25,
		Gender:            "male",
		Diet_Level:        1,
		Favorite_Recipes:  []int32{1, 2, 3},
		Address:           "myadress",
		Role:              "user",
		Preferred_Meals:   []string{"breakfast"},
		Likes:             []string{"kebap"},
		Dislikes:          []string{"onion"},
		Prohibits:         []string{"sugar"},
		Wants:             `gain`,
		Favorite_Cuisines: []string{"italian"},
	}
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&User{})
	return db, user
}
func Destruct(db *gorm.DB) {
	db.Exec("DROP TABLE users")
}
func TestIsAuth(t *testing.T) {
	db, user := Construct()
	user.Signup(db)
	user2 := user
	user2.Password = "wrongpass"
	tests := []struct {
		input  User
		output bool
		err    error
	}{
		{input: user, output: true, err: nil},
		{input: user2, output: false, err: nil},
	}
	for _, test := range tests {
		res, err := test.input.IsAuth(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if res != test.output {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}
	}
	Destruct(db)
}
func TestHashPassword(t *testing.T) {
	db, user := Construct()
	tests := []struct {
		input string
		err   error
	}{
		{input: "secret", err: nil},
	}
	for _, test := range tests {

		_, err := user.HashPassword(test.input)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(db)
}
func TestCheckPasswordHash(t *testing.T) {
	db, user := Construct()
	user.Password, _ = user.HashPassword(user.Password)
	user.Create(db)
	tests := []struct {
		input  string
		output bool
	}{
		{input: "secret", output: false},
		{input: "testpass", output: true},
	}
	for _, test := range tests {
		out := user.CheckPasswordHash(test.input, user.Password)
		if test.output != out {
			t.Errorf("Result is: %v . Expected: %v", out, test.output)
		}
	}
	Destruct(db)
}
func TestSignup(t *testing.T) {
	db, user := Construct()
	tests := []struct {
		input User
		err   error
	}{
		{input: user, err: nil},
	}
	for _, test := range tests {
		err := test.input.Signup(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(db)
}
func TestIsCredentialsExist(t *testing.T) {
	db, user := Construct()
	user.Create(db)
	user2 := User{Username: "notexisteduser", Email: "notexist@test.com"}
	tests := []struct {
		input User
		err   error
	}{
		{input: user2, err: nil},
		//{input: user2, err: nil},
	}
	for _, test := range tests {
		err := test.input.IsCredentialsExist(db)
		if err != test.err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(db)
}
