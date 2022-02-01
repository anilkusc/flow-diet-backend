package recipe

import (
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model              `json:"-" swaggerignore:"true"`
	Name                    string
	Ingredients             []ingredient.Ingredient `json:"ingredients" swaggerignore:"true" example:"{will be implemented}"`
	Ingredients_String      string                  `gorm:"not null" json:"-" `
	Preperation             string                  `gorm:"not null" json:"preperation" example:"bla bla bla"`
	Preperation_Time_minute uint16                  `gorm:"not null" json:"preperation_time" example:"15"`
	Cooking_Time_Minute     uint16                  `gorm:"not null" json:"cooking_time_minute" example:"10"`
	Calori                  uint16                  `gorm:"not null" json:"calori" example:"252"`
	Photo_Urls              string                  `gorm:"not null" json:"photo_urls" example:"[{'url':'S3URL'}]"`
	Video_Urls              string                  `gorm:"not null" json:"video_urls" example:"[{'url':'S3URL'}]"`
	For_How_Many_People     uint8                   `gorm:"not null" json:"for_how_many_people" example:"2"`
}
