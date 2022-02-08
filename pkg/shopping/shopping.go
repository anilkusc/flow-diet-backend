package shopping

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Shopping struct {
	gorm.Model  `swaggerignore:"true"`
	Ingredients pq.Int32Array `json:"ingredients" gorm:"type:int[]" example:"1,2,3"` //Ingredients ID
	User_Id     uint          `gorm:"not null" json:"user_id" example:"1"`
	Start_Date  uint          `gorm:"not null" json:"start_date" example:"1643743444"`
	End_Date    uint          `gorm:"not null" json:"end_date" example:"1643743448"`
}
