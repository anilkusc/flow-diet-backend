package app

import (
	"encoding/json"
	"reflect"
	"testing"

	user "github.com/anilkusc/flow-diet-backend/pkg/user"
)

func TestSignup(t *testing.T) {
	app, user := Construct()
	userJson, _ := json.Marshal(user)
	tests := []struct {
		input string
		err   error
	}{
		{input: string(userJson), err: nil},
	}
	for _, test := range tests {
		err := app.Signup(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct(app)
}
func TestSignin(t *testing.T) {
	app, usr := Construct()

	userJson, _ := json.Marshal(usr)
	app.Signup(string(userJson))
	usr.Password = ""

	tests := []struct {
		input  string
		output user.User
		isAuth bool
		err    error
	}{
		{input: string(userJson), output: usr, isAuth: true, err: nil},
	}
	for _, test := range tests {
		output, isAuth, err := app.Signin(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if isAuth != test.isAuth {
			t.Errorf("Result isAuth is: %v . Expected: %v", isAuth, test.isAuth)
		}
		output.Favorite_Recipes_String, _ = output.ArrayToJson(output.Favorite_Recipes)
		if !reflect.DeepEqual(output, test.output) {
			t.Errorf("Result is: %v . Expected: %v", output, test.output)
		}

	}
	Destruct(app)
}
