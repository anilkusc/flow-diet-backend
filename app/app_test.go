package app

import (
	"time"

	"github.com/anilkusc/flow-diet-backend/pkg/calendar"
	user "github.com/anilkusc/flow-diet-backend/pkg/user"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Construct() (App, user.User, calendar.Calendar) {
	godotenv.Load("../.env")
	app := App{}
	app.Init()
	user := user.User{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Username:                "testuser1",
		Name:                    "test user",
		Password:                "testpass",
		Weight:                  70,
		Height:                  173,
		Age:                     25,
		Diet:                    "omnivor",
		Favorite_Recipes:        []uint{1, 2, 3},
		Favorite_Recipes_String: "[1,2,3]",
		Address:                 "",
		Role:                    "user",
	}
	var calendar = calendar.Calendar{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Recipe_Id:  1,
		User_Id:    1,
		Date_Epoch: 1643743444,
	}
	return app, user, calendar
}
func Destruct(app App) {
	app.DB.Exec("DROP TABLE users")
}
