package recipe

import (
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Construct() (*gorm.DB, Recipe) {

	godotenv.Load("../../.env")
	var recipe = Recipe{
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
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Recipe{})
	return db, recipe
}
func Destruct(db *gorm.DB) {
	db.Exec("DROP TABLE recipes")
}
