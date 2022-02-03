package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/anilkusc/flow-diet-backend/pkg/calendar"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/material"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/measurement"
	"github.com/anilkusc/flow-diet-backend/pkg/shopping"
	user "github.com/anilkusc/flow-diet-backend/pkg/user"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Construct() (App, string, user.User, calendar.Calendar, recipe.Recipe, shopping.Shopping) {
	godotenv.Load("../.env")
	app := App{}
	app.Init()
	usr := user.User{
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
		Role:                    "admin", //"user"
	}
	userJson, _ := json.Marshal(usr)
	app.Signup(string(userJson))

	req, _ := http.NewRequest("POST", "/user/signin", strings.NewReader(string(userJson)))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.SigninHandler)

	handler.ServeHTTP(rr, req)
	sessionCookie := rr.Header()["Set-Cookie"][0]
	ck := strings.Split(sessionCookie, " ")
	ck = strings.Split(ck[0], "session=")
	cookie := ck[1]
	signInCookie := cookie[:len(cookie)-1]
	app.DB.Exec("DROP TABLE users")
	app.DB.AutoMigrate(&user.User{})
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
	var shopping = shopping.Shopping{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
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
		Ingredients_String: `[{"measurement":{"size":200,"quantity":"gram"},"material":{"type":"fruit","name":"banana","photo_urls":"['s3link1','s3link2']"},"isexist":false,"isoptional":false}]`,
		Start_Date:         "1643743444",
		End_Date:           "1643743448",
		User_Id:            1,
	}
	return app, signInCookie, usr, calendar, recipe, shopping
}
func Destruct(app App) {
	app.DB.Exec("DROP TABLE users")
	app.DB.Exec("DROP TABLE calendars")
	app.DB.Exec("DROP TABLE recipes")
	app.DB.Exec("DROP TABLE shoppings")
}
