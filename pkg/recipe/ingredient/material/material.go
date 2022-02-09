package material

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Material struct {
	gorm.Model          `swaggerignore:"true"`
	Material_Name       string         `gorm:"not null" json:"material_name" example:"bread"`
	Size                float32        `gorm:"not null" json:"measurement_size" example:"200"`
	Quantity            string         `gorm:"not null" json:"measurement_quantity" example:"gram"`
	Material_Photo_Urls pq.StringArray `gorm:"type:text[]" json:"material_photo_urls" example:"S3URL1,S3URL2"`
	Tags                pq.StringArray `gorm:"type:text[]" json:"material_tags" example:"vegan"`
}
