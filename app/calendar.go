package app

import (
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
