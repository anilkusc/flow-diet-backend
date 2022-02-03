package app

import (
	"encoding/json"

	"github.com/anilkusc/flow-diet-backend/pkg/shopping"
)

func (app *App) ListShoppings(userid uint) (string, error) {
	var shoppings []shopping.Shopping
	var shopping shopping.Shopping
	var err error
	shopping.User_Id = userid
	shoppings, err = shopping.List(app.DB)
	if err != nil {
		return "", err
	}
	shoppingsList, err := json.Marshal(shoppings)
	if err != nil {
		return "", err
	}
	return string(shoppingsList), nil
}

func (app *App) CreateShopping(shoppingJson string) error {
	var shopping shopping.Shopping
	var err error
	err = json.Unmarshal([]byte(shoppingJson), &shopping)
	if err != nil {
		return err
	}

	err = shopping.Create(app.DB)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) GetShopping(shoppingJson string) (string, error) {

	var shopping shopping.Shopping
	var err error
	err = json.Unmarshal([]byte(shoppingJson), &shopping)
	if err != nil {
		return "", err
	}

	shopping.Read(app.DB)
	if err != nil {
		return "", err
	}
	shoppingStr, err := json.Marshal(shopping)
	if err != nil {
		return "", err
	}
	return string(shoppingStr), nil
}

func (app *App) UpdateShopping(shoppingJson string) error {

	var shopping shopping.Shopping
	err := json.Unmarshal([]byte(shoppingJson), &shopping)
	if err != nil {
		return err
	}

	shopping.Update(app.DB)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) DeleteShopping(shoppingJson string) error {

	var shopping shopping.Shopping
	err := json.Unmarshal([]byte(shoppingJson), &shopping)
	if err != nil {
		return err
	}
	err = shopping.Read(app.DB)
	if err != nil {
		return err
	}
	shopping.Delete(app.DB)
	if err != nil {
		return err
	}

	return nil
}
