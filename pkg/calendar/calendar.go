package calendar

import (
	"gorm.io/gorm"
)

type Calendar struct {
	gorm.Model `swaggerignore:"true"`
	Recipe_Id  uint   `gorm:"not null" json:"recipe_id" example:"1"`
	User_Id    uint   `gorm:"not null" json:"user_id" example:"1"`
	Meal       string `gorm:"not null" json:"meal" example:"breakfast"`
	Date_Epoch uint64 `gorm:"not null" json:"date_epoch" example:"1643743444"`
	Prepared   bool   `gorm:"not null" json:"prepared" example:"false"`
}
