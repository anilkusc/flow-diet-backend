package calendar

import (
	"gorm.io/gorm"
)

type Calendar struct {
	gorm.Model `json:"-" swaggerignore:"true"`
	Recipe_Id  uint   `gorm:"not null" json:"recipe_id" example:"1"`
	User_Id    uint   `gorm:"not null" json:"user_id" example:"1"`
	Date_Epoch string `gorm:"not null" json:"date_epoch" example:"1643743444"`
}
