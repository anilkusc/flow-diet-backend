package ingredient

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model           `json:"-" swaggerignore:"true"`
	Measurement_Size     float32        `gorm:"not null" json:"measurement_size" example:"200"`
	Measurement_Quantity string         `gorm:"not null" json:"measurement_quantity" example:"gram"`
	Material_Photo_Urls  pq.StringArray `gorm:"type:text[]" json:"material_photo_urls" example:"S3URL1,S3URL2"`
	Material_Name        string         `gorm:"not null" json:"material_name" example:"bread"`
	IsExist              bool           `json:"isexist" example:"false"`
	IsOptional           bool           `json:"isoptional" example:"true"`
}
