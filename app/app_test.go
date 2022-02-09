package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/anilkusc/flow-diet-backend/pkg/calendar"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient/material"
	"github.com/anilkusc/flow-diet-backend/pkg/recommendation"
	"github.com/anilkusc/flow-diet-backend/pkg/search"
	"github.com/anilkusc/flow-diet-backend/pkg/shopping"
	user "github.com/anilkusc/flow-diet-backend/pkg/user"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Construct() (App, string, user.User, calendar.Calendar, recipe.Recipe, shopping.Shopping, search.Search, recommendation.Recommendation, material.Material) {

	godotenv.Load("../.env")
	app := App{}
	app.Init()
	usr := user.User{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Username:          "testuser1",
		Name:              "test user",
		Email:             "testmail@test.com",
		Phone:             "+905355353535",
		Password:          "testpass",
		Weight:            70,
		Height:            173,
		Age:               25,
		Gender:            "male",
		Diet_Level:        1,
		Favorite_Recipes:  []int32{1, 2, 3},
		Address:           "myadress",
		Role:              "admin",
		Preferred_Meals:   []string{"breakfast"},
		Likes:             []string{"kebap"},
		Dislikes:          []string{"onion"},
		Prohibits:         []string{"sugar"},
		Wants:             `gain`,
		Favorite_Cuisines: []string{"italian"},
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
	calendar := calendar.Calendar{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Recipe_Id:  1,
		User_Id:    1,
		Meal:       "breakfast",
		Date_Epoch: 1643743444,
		Prepared:   false,
	}

	recipe := recipe.Recipe{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Title:                   "Test Recipe",
		Ingredients:             []int32{1, 2},
		Preperation:             []string{"blalblabla"},
		Preperation_Time_minute: 15,
		Cooking_Time_Minute:     15,
		Calori:                  255,
		Recipe_Diet_Level:       1,
		Photo_Urls:              []string{"S3URL1", "S3URL2"},
		Video_Urls:              []string{"S3URL1", "S3URL2"},
		For_How_Many_People:     2,
		Tags:                    []string{"vegan", "kebap", "cola"},
		Appropriate_Meals:       []string{"breakfast", "afternoon"},
		Cuisines:                []string{"italian"},
	}
	var shopping = shopping.Shopping{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Ingredients: []int32{1, 2, 3},
		Start_Date:  1643743444,
		End_Date:    1643743448,
		User_Id:     1,
	}
	search := search.Search{Word: "some"}

	var recommendation = recommendation.Recommendation{
		Users_Preferred_Meals: []string{"breakfast", "noon", "night"},
		Users_Prohibits:       []string{"sugar"},
		Users_Diet_Level:      1,
		Users_Dislikes:        []string{"onion", "tomato"},
		Users_Likes:           []string{"chicken", "fish"},
		Users_Cousines:        []string{"italian", "mediterrian"},

		Recipe_IDsAppropriateMeals: map[uint][]string{1: {"breakfast", "night"}, 2: {"noon"}, 3: {"snack"}},
		Recipe_IDsTags:             map[uint][]string{1: {"sugar", "tea"}, 2: {"fish", "chips"}, 3: {"rice", "sushi"}, 4: {"vegaterian", "egg"}},
		Recipe_IDsCousines:         map[uint][]string{1: {"mediterrian", "russia"}, 2: {"china", "japan"}, 3: {"italian", "spanish"}, 4: {"ireland"}},
		Recipe_IDsDietlevel:        map[uint]uint{1: 1, 2: 2, 3: 2, 4: 1},
		All_Recipes_IDs:            []uint{1, 2, 3, 4},
		Recipe_IDsPoints:           map[uint]uint{},

		Meal_Factor:         2,
		Like_Factor:         3,
		Dislike_Factor:      2,
		Cousine_Factor:      2,
		Recommended_Recipes: []uint{}, // it is sorted by recommended points.
	}
	var material = material.Material{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Material_Name:       "bread",
		Material_Photo_Urls: []string{"S3URL1", "S3URL2"},
		Tags:                []string{"vegan"},
		Size:                200,
		Quantity:            "gram",
	}
	return app, signInCookie, usr, calendar, recipe, shopping, search, recommendation, material
}
func Destruct(app App) {
	app.DB.Exec("DROP TABLE users")
	app.DB.Exec("DROP TABLE calendars")
	app.DB.Exec("DROP TABLE recipes")
	app.DB.Exec("DROP TABLE shoppings")
	app.DB.Exec("DROP TABLE materials")
}
