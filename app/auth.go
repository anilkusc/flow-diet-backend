package app

import (
	"encoding/json"

	user "github.com/anilkusc/flow-diet-backend/pkg/user"
)

func (app *App) Signup(userJson string) error {
	var user user.User
	err := json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return err
	}
	err = user.Signup(app.DB)
	if err != nil {
		return err
	}
	return nil
}
func (app *App) Signin(userJson string) (bool, error) {
	var user user.User
	var isauth bool
	err := json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return false, err
	}
	isauth, err = user.IsAuth(app.DB)
	if err != nil {
		return false, err
	}
	return isauth, nil
}
