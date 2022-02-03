package shopping

import (
	"encoding/json"

	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient"
	"gorm.io/gorm"
)

type Shopping struct {
	gorm.Model         `swaggerignore:"true"`
	Ingredients        []ingredient.Ingredient `gorm:"-" json:"ingredients"`
	Ingredients_String string                  `json:"-" swaggerignore:"true"`
	User_Id            uint                    `gorm:"not null" json:"user_id" example:"1"`
	Start_Date         string                  `gorm:"not null" json:"start_date" example:"1643743444"`
	End_Date           string                  `gorm:"not null" json:"end_date" example:"1643743448"`
}

func (s *Shopping) ArrayToJson(arr []ingredient.Ingredient) (string, error) {

	shoppingString, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}
	return string(shoppingString), nil
}

func (s *Shopping) JsonToArray(arr string) ([]ingredient.Ingredient, error) {

	var array []ingredient.Ingredient
	err := json.Unmarshal([]byte(arr), &array)
	if err != nil {
		return array, err
	}
	return array, nil
}
