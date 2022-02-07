package app

import (
	"encoding/json"
	"errors"

	"github.com/anilkusc/flow-diet-backend/pkg/calendar"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe"
	"github.com/anilkusc/flow-diet-backend/pkg/shopping"
)

func (app *App) GetMyCalendar(userid uint) ([]calendar.Calendar, error) {
	var calendars []calendar.Calendar
	var calendar calendar.Calendar
	var err error
	calendar.User_Id = userid

	calendars, err = calendar.List(app.DB)
	if err != nil {
		return calendars, err
	}
	return calendars, nil
}

func (app *App) CreateCalendar(calendarsJson string, userid uint) error {
	clndr := calendar.Calendar{}
	var calendars []calendar.Calendar
	var recipe recipe.Recipe
	err := json.Unmarshal([]byte(calendarsJson), &calendars)
	if err != nil {
		return err
	}
	// TODO this type of controls should be migrated to middlewares.
	if len(calendars) < 1 && len(calendars) > 50 {
		return errors.New("object length must not be bigger than 50 and smaller than 1")
	}

	shoppingList := shopping.Shopping{User_Id: userid}
	minDate := calendars[0].Date_Epoch
	maxDate := calendars[0].Date_Epoch

	for i := range calendars {
		if calendars[i].Date_Epoch < minDate {
			minDate = calendars[i].Date_Epoch
		}
		if calendars[i].Date_Epoch > maxDate {
			maxDate = calendars[i].Date_Epoch
		}

		recipe.ID = calendars[i].Recipe_Id
		recipe.Read(app.DB)
		for j := range recipe.Ingredients {
			shoppingList.Ingredients = append(shoppingList.Ingredients, recipe.Ingredients[j])
		}
		calendars[i].User_Id = userid
	}

	shoppingList.Start_Date = uint(minDate)
	shoppingList.End_Date = uint(maxDate)
	err = shoppingList.Create(app.DB)
	if err != nil {
		return err
	}

	err = clndr.CreateBulk(app.DB, calendars)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) UpdateCalendar(calendarJson string, userid uint) error {

	var calendar calendar.Calendar
	err := json.Unmarshal([]byte(calendarJson), &calendar)
	if err != nil {
		return err
	}
	if calendar.ID == 0 {
		return errors.New("id is not specified for update request")
	}

	calendar.User_Id = userid
	err = calendar.Update(app.DB)

	return err
}

func (app *App) DeleteCalendar(calendarJson string, userid uint) error {

	var calendar calendar.Calendar
	err := json.Unmarshal([]byte(calendarJson), &calendar)
	if err != nil {
		return err
	}
	if calendar.ID == 0 {
		return errors.New("id is not specified in the request")
	}

	err = calendar.Read(app.DB)
	if err != nil {
		return err
	}

	if calendar.User_Id != userid {
		return errors.New("this user is not allowed to delete this recipe in the calendar")
	}

	err = calendar.Delete(app.DB)

	return err
}
