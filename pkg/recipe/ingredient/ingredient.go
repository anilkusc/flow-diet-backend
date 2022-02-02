package ingredient

import (
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/material"
	"github.com/anilkusc/flow-diet-backend/pkg/recipe/measurement"
)

type Ingredient struct {
	//gorm.Model  `json:"-" swaggerignore:"true"`
	//Size        float32
	Measurement measurement.Measurement `json:"measurement" example:"1"`
	Material    material.Material       `json:"material" example:"1"`
	IsExist     bool                    `json:"isexist" example:"false"`
	IsOptional  bool                    `json:"isoptional" example:"true"`
}
