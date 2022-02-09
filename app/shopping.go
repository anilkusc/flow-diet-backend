package app

import (
	"encoding/json"
	"errors"
	"time"

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
func (app *App) ListShoppingsWithDateInterval(userid uint, shoppingJson string) (string, error) {

	//TODO: Convert every dates to time.Time format

	var shoppings []shopping.Shopping
	var shopping shopping.Shopping
	var err error

	err = json.Unmarshal([]byte(shoppingJson), &shopping)
	if err != nil {
		return "", err
	}
	// 5000000 approximately 3 months
	if (time.Now().Unix()-int64(shopping.Start_Date)) > 5000000 || (time.Now().Unix()-int64(shopping.Start_Date)) < 0 {
		return "", errors.New("cannot be queried for more than 3 months")
	}
	// 600000 approximately 1 week
	if (shopping.End_Date - shopping.Start_Date) > 600000 {
		return "", errors.New("time interval cannot be more than 1 week")
	}

	shopping.User_Id = userid
	shoppings, err = shopping.ListByDateInterval(app.DB)
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
