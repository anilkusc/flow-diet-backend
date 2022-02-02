package app

import (
	"encoding/json"
	"errors"

	"github.com/anilkusc/flow-diet-backend/pkg/calendar"
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

func (app *App) CreateCalendar(calendarJson string, userid uint) error {

	var calendar calendar.Calendar
	err := json.Unmarshal([]byte(calendarJson), &calendar)
	if err != nil {
		return err
	}
	calendar.User_Id = userid
	err = calendar.Create(app.DB)

	return err
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
