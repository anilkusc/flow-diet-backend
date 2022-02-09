package ingredient

import (
	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model  `swaggerignore:"true"`
	Material_Id uint `gorm:"not null" json:"material_id" example:"1"`
	IsOptional  bool `gorm:"not null" json:"isoptional" example:"true"`
}
