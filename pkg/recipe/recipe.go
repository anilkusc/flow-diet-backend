package recipe

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model              `swaggerignore:"true"`
	Title                   string         `gorm:"not null" json:"title" example:"Sushi With Wassabi"`
	Recipe_Diet_Level       uint           `gorm:"not null" json:"recipe_diet_level" example:"1"`
	Tags                    pq.StringArray `gorm:"type:text[]" json:"tags" example:"vegan"`
	Ingredients             pq.Int32Array  `gorm:"not null;type:int[]"  json:"ingredients" example:"1,2" `
	Preperation             pq.StringArray `gorm:"type:text[]" json:"preperation" example:"bla bla bla"`
	Preperation_Time_minute uint16         `gorm:"not null" json:"preperation_time" example:"15"`
	Cooking_Time_Minute     uint16         `gorm:"not null" json:"cooking_time_minute" example:"10"`
	Calori                  uint16         `json:"calori" example:"252"`
	Photo_Urls              pq.StringArray `gorm:"type:text[]" json:"photo_urls" example:"S3URL1,S3URL2"`
	Video_Urls              pq.StringArray `gorm:"type:text[]" json:"video_urls" example:"S3URL1,S3URL2"`
	For_How_Many_People     uint8          `gorm:"not null" json:"for_how_many_people" example:"2"`
	Appropriate_Meals       pq.StringArray `gorm:"type:text[]" json:"appropriate_meals" example:"breakfast,snack"` // Breakfast, Snack , Noon , AfterNoon , Evening , Night // It should be added by appropriate wieght sort.
	Cuisines                pq.StringArray `gorm:"type:text[]" json:"cousines" example:"italian"`
}
