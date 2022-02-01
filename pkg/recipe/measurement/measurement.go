package measurement

import "gorm.io/gorm"

type Measurement struct {
	gorm.Model `json:"-" swaggerignore:"true"`
	Size       float32 `gorm:"not null" json:"size" example:"2"`
	Quantity   string  `gorm:"not null" json:"quantity" example:"gram"`
}
