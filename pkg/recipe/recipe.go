package recipe

import (
	"encoding/json"

	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model              //`json:"-" swaggerignore:"true"`
	Name                    string
	Ingredients             []ingredient.Ingredient `gorm:"-" json:"ingredients" swaggerignore:"true" example:"[{'measurement':{'size':200,'quantity':'gram'},'material':{'type':'fruit','name':'banana','photo_urls':'['s3link1','s3link2']'},'isexist':false,'isoptional':false}]"`
	Ingredients_String      string                  `gorm:"not null" json:"-" `
	Preperation             string                  `gorm:"not null" json:"preperation" example:"bla bla bla"`
	Preperation_Time_minute uint16                  `gorm:"not null" json:"preperation_time" example:"15"`
	Cooking_Time_Minute     uint16                  `gorm:"not null" json:"cooking_time_minute" example:"10"`
	Calori                  uint16                  `gorm:"not null" json:"calori" example:"252"`
	Photo_Urls              string                  `gorm:"not null" json:"photo_urls" example:"[{'url':'S3URL'}]"`
	Video_Urls              string                  `gorm:"not null" json:"video_urls" example:"[{'url':'S3URL'}]"`
	For_How_Many_People     uint8                   `gorm:"not null" json:"for_how_many_people" example:"2"`
}

func (r *Recipe) ArrayToJson(arr []ingredient.Ingredient) (string, error) {

	recipeString, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}
	return string(recipeString), nil
}

func (r *Recipe) JsonToArray(arr string) ([]ingredient.Ingredient, error) {

	var array []ingredient.Ingredient
	err := json.Unmarshal([]byte(arr), &array)
	if err != nil {
		return array, err
	}
	return array, nil
}
