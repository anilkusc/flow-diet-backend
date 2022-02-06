package recipe

import (
	"encoding/json"

	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model               `swaggerignore:"true"`
	Title                    string                  `gorm:"not null" json:"title" example:"Sushi With Wassabi"`
	Recipe_Diet_Level        uint                    `gorm:"not null" json:"recipe_diet_level" example:"1"`
	Tags                     []string                `gorm:"-" json:"tags" example:"vegan"`
	Tags_String              string                  `gorm:"not null" json:"-" swaggerignore:"true"`
	Ingredients              []ingredient.Ingredient `gorm:"-" json:"ingredients" `
	Ingredients_String       string                  `gorm:"not null" json:"-" swaggerignore:"true" `
	Preperation              string                  `gorm:"not null" json:"preperation" example:"bla bla bla"`
	Preperation_Time_minute  uint16                  `gorm:"not null" json:"preperation_time" example:"15"`
	Cooking_Time_Minute      uint16                  `gorm:"not null" json:"cooking_time_minute" example:"10"`
	Calori                   uint16                  ` json:"calori" example:"252"`
	Photo_Urls               []string                `gorm:"-" json:"photo_urls" example:"S3URL1,S3URL2"`
	Video_Urls               []string                `gorm:"-" json:"video_urls" example:"S3URL1,S3URL2"`
	Photo_Urls_String        string                  `json:"-" swaggerignore:"true"`
	Video_Urls_String        string                  `json:"-" swaggerignore:"true"`
	For_How_Many_People      uint8                   `gorm:"not null" json:"for_how_many_people" example:"2"`
	Appropriate_Meals        []string                `gorm:"-" json:"appropriate_meals" example:"breakfast,snack"` // Breakfast, Snack , Noon , AfterNoon , Evening , Night // It should be added by appropriate wieght sort.
	Appropriate_Meals_String string                  `json:"-" swaggerignore:"true"`
	Cousines                 []string                `gorm:"-" json:"cousines" example:"italian"`
	Cousines_String          string                  `json:"-" swaggerignore:"true"`
}

func (r *Recipe) IngredientToJson(arr []ingredient.Ingredient) (string, error) {

	recipeString, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}
	return string(recipeString), nil
}

func (r *Recipe) JsonToIngredient(arr string) ([]ingredient.Ingredient, error) {

	var array []ingredient.Ingredient
	err := json.Unmarshal([]byte(arr), &array)
	if err != nil {
		return array, err
	}
	return array, nil
}
func (r *Recipe) ArrayToJson(arr []string) (string, error) {

	recipeString, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}
	return string(recipeString), nil
}

func (r *Recipe) JsonToArray(arr string) ([]string, error) {

	var array []string
	err := json.Unmarshal([]byte(arr), &array)
	if err != nil {
		return array, err
	}
	return array, nil
}
