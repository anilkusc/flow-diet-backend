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
	"github.com/anilkusc/flow-diet-backend/pkg/recommendation"
	"github.com/anilkusc/flow-diet-backend/pkg/search"
	"github.com/anilkusc/flow-diet-backend/pkg/shopping"
	user "github.com/anilkusc/flow-diet-backend/pkg/user"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Construct() (App, string, user.User, calendar.Calendar, recipe.Recipe, shopping.Shopping, search.Search, recommendation.Recommendation) {
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
		Email:                   "testmail@testail.com",
		Phone:                   "+905355556789",
		Weight:                  70,
		Height:                  173,
		Age:                     25,
		Gender:                  "male",
		Diet_Level:              1,
		Favorite_Recipes:        []uint{1, 2, 3},
		Favorite_Recipes_String: "[1,2,3]",
		Address:                 "my address 123",
		Role:                    "admin", //"user"
		Preferred_Meals:         []string{"breakfast"},
		Preferred_Meals_String:  `["breakfast"]`,
		Likes:                   []string{"kebap"},
		Likes_String:            `["kebap"]`,
		Dislikes:                []string{"onion"},
		Dislikes_String:         `["onion"]`,
		Prohibits:               []string{"sugar"},
		Prohibits_String:        `["sugar"]`,
		Wants:                   `gain`,
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
		Title: "Test Recipe",
		Ingredients: []ingredient.Ingredient{
			{
				Measurement: measurement.Measurement{
					Size:     200,
					Quantity: "gram",
				},
				Material: material.Material{
					Name:                "banana",
					Material_Photo_Urls: []string{"S3URL1", "S3URL2"},
				},
				IsExist:    false,
				IsOptional: false,
			},
		},
		Ingredients_String:       `[{"measurement":{"size":200,"quantity":"gram"},"material":{"name":"banana","material_photo_urls":["S3URL1","S3URL2"]},"isexist":false,"isoptional":false}]`,
		Preperation:              "Cook the chickens!",
		Preperation_Time_minute:  15,
		Cooking_Time_Minute:      15,
		Calori:                   255,
		Photo_Urls:               []string{"S3URL1", "S3URL2"},
		Video_Urls:               []string{"S3URL1", "S3URL2"},
		Photo_Urls_String:        `["S3URL1", "S3URL2"]`,
		Video_Urls_String:        `["S3URL1", "S3URL2"]`,
		For_How_Many_People:      2,
		Tags:                     []string{"vegan", "kebap", "cola"},
		Tags_String:              `["vegan", "kebap", "cola"]`,
		Appropriate_Meals:        []string{"breakfast", "afternoon"},
		Appropriate_Meals_String: `["breakfast","afternoon"]`,
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
					Name:                "banana",
					Material_Photo_Urls: []string{"S3URL1", "S3URL2"},
				},
				IsExist:    false,
				IsOptional: false,
			},
		},
		Ingredients_String: `[{"measurement":{"size":200,"quantity":"gram"},"material":{"name":"banana","material_photo_urls":["S3URL1","S3URL2"]},"isexist":false,"isoptional":false}]`,
		Start_Date:         "1643743444",
		End_Date:           "1643743448",
		User_Id:            1,
	}
	search := search.Search{Word: "some"}

	var recommendation = recommendation.Recommendation{
		Users_Preferred_Meals: []string{"breakfast", "noon", "night"},
		Users_Prohibits:       []string{"sugar"},
		Users_Diet_Level:      1,
		Users_Dislikes:        []string{"onion", "tomato"},
		Users_Likes:           []string{"chicken", "fish"},

		Recipe_IDsAppropriateMeals: map[uint][]string{1: {"breakfast", "night"}, 2: {"noon"}, 3: {"snack"}},
		Recipe_IDsTags:             map[uint][]string{1: {"sugar", "tea"}, 2: {"fish", "chips"}, 3: {"rice", "sushi"}, 4: {"vegaterian", "egg"}},
		Recipe_IDsDietlevel:        map[uint]uint{1: 1, 2: 2, 3: 2, 4: 1},
		All_Recipes_IDs:            []uint{1, 2, 3, 4},
		Recipe_IDsPoints:           map[uint]uint{},

		Meal_Factor:         2,
		Like_Factor:         3,
		Dislike_Factor:      2,
		Recommended_Recipes: []uint{}, // it is sorted by recommended points.
	}

	return app, signInCookie, usr, calendar, recipe, shopping, search, recommendation
}
func Destruct(app App) {
	app.DB.Exec("DROP TABLE users")
	app.DB.Exec("DROP TABLE calendars")
	app.DB.Exec("DROP TABLE recipes")
	app.DB.Exec("DROP TABLE shoppings")
}
