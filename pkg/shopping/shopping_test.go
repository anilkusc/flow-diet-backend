package shopping

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

func Construct() (*gorm.DB, Shopping) {
	godotenv.Load("../../.env")
	var shopping = Shopping{
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
					Tags:                []string{"vegan", "fruit"},
					Material_Photo_Urls: []string{"S3URL1", "S3URL2"},
				},
				IsExist:    false,
				IsOptional: false,
			},
		},
		Ingredients_String: `[{"measurement":{"size":200,"quantity":"gram"},"material":{"name":"banana","tags":["vegan","fruit"],"material_diet_level":0,"material_photo_urls":["S3URL1","S3URL2"]},"isexist":false,"isoptional":false}]`,
		Start_Date:         "1643743444",
		End_Date:           "1643743448",
		User_Id:            1,
	}
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Shopping{})
	return db, shopping
}
func Destruct(db *gorm.DB) {
	db.Exec("DROP TABLE shoppings")
}

func TestArrayToJson(t *testing.T) {
	db, shopping := Construct()

	tests := []struct {
		input  []ingredient.Ingredient
		output string
		err    error
	}{
		{input: shopping.Ingredients, output: shopping.Ingredients_String, err: nil},
	}
	for _, test := range tests {
		res, err := shopping.ArrayToJson(test.input)
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
	db, shopping := Construct()

	tests := []struct {
		input  string
		output []ingredient.Ingredient
		err    error
	}{
		{input: shopping.Ingredients_String, output: shopping.Ingredients, err: nil},
	}
	for _, test := range tests {
		res, err := shopping.JsonToArray(test.input)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if !reflect.DeepEqual(test.output, res) {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}
	}
	Destruct(db)
}
