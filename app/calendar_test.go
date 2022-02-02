package app

import (
	"reflect"
	"testing"
	"time"

	"github.com/anilkusc/flow-diet-backend/pkg/calendar"
	"gorm.io/gorm"
)

func TestGetMyCalendar(t *testing.T) {
	app, _, clndr := Construct()
	clndr.Create(app.DB)
	tests := []struct {
		input  uint
		output []calendar.Calendar
		err    error
	}{
		{input: 1, output: []calendar.Calendar{clndr}, err: nil},
	}
	for _, test := range tests {
		res, err := app.GetMyCalendar(test.input)
		if err != nil {
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
	Destruct(app)
}
