package ingredient

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model `json:"-" swaggerignore:"true"`
	Materials  pq.Int32Array `gorm:"not null;type:int[]" json:"materials" example:"1,2,3"`
	IsOptional bool          `gorm:"not null" json:"isoptional" example:"true"`
}
