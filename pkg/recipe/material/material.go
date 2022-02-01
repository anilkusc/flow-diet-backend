package material

import "gorm.io/gorm"

type Material struct {
	gorm.Model `json:"-" swaggerignore:"true"`
	Type       string `gorm:"not null" json:"type" example:"fruit"`
	Name       string `gorm:"not null" json:"name" example:"banana"`
	Photo_Urls string `json:"photo_urls" example:"[{'url':'exampleS3URL'}]"`
}
