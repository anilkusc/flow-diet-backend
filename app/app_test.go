package app

import (
	"time"

	"github.com/anilkusc/flow-diet-backend/pkg/calendar"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/material"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/measurement"
	user "github.com/anilkusc/flow-diet-backend/pkg/user"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Construct() (App, user.User, calendar.Calendar, recipe.Recipe) {
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
	var recipe = recipe.Recipe{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Name: "Test Recipe",
		Ingredients: []ingredient.Ingredient{
			{
				Measurement: measurement.Measurement{
					Size:     200,
					Quantity: "gram",
				},
				Material: material.Material{
					Type:       "fruit",
					Name:       "banana",
					Photo_Urls: "['s3link1','s3link2']",
				},
				IsExist:    false,
				IsOptional: false,
			},
		},
		Ingredients_String:      `[{"measurement":{"size":200,"quantity":"gram"},"material":{"type":"fruit","name":"banana","photo_urls":"['s3link1','s3link2']"},"isexist":false,"isoptional":false}]`,
		Preperation:             "Cook the chickens!",
		Preperation_Time_minute: 15,
		Cooking_Time_Minute:     15,
		Calori:                  255,
		Photo_Urls:              "['S3URL1','S3URL2']",
		Video_Urls:              "['S3URL1','S3URL2']",
		For_How_Many_People:     2,
	}
	return app, user, calendar, recipe
}
func Destruct(app App) {
	app.DB.Exec("DROP TABLE users")
	app.DB.Exec("DROP TABLE calendars")
	app.DB.Exec("DROP TABLE recipes")
}
