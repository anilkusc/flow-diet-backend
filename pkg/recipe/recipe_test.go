package recipe

import (
	"reflect"
	"testing"
	"time"

	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/material"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/measurement"
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
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Recipe{})
	return db, recipe
}
func Destruct(db *gorm.DB) {
	db.Exec("DROP TABLE recipes")
}

func TestArrayToJson(t *testing.T) {
	db, recipe := Construct()

	tests := []struct {
		input  []ingredient.Ingredient
		output string
		err    error
	}{
		{input: recipe.Ingredients, output: recipe.Ingredients_String, err: nil},
	}
	for _, test := range tests {
		res, err := recipe.ArrayToJson(test.input)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if test.output != res {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}
	}
	Destruct(db)
}
func TestJsonToArray(t *testing.T) {
	db, recipe := Construct()

	tests := []struct {
		input  string
		output []ingredient.Ingredient
		err    error
	}{
		{input: recipe.Ingredients_String, output: recipe.Ingredients, err: nil},
	}
	for _, test := range tests {
		res, err := recipe.JsonToArray(test.input)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if !reflect.DeepEqual(test.output, res) {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}
	}
	Destruct(db)
}
