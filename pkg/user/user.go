package user

import (
	"encoding/json"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username                string `gorm:"unique;not null"`
	Name                    string
	Password                string
	Weight                  uint8
	Height                  uint8
	Age                     uint8
	Diet                    string //vegaterian , vegan , omnivor , carnivor
	Favorite_Recipes        []uint `gorm:"-"` //
	Favorite_Recipes_String string
	Address                 string
	Role                    string // user,admin,dietician,editor
}

func (u *User) ArrayToJson(arr []uint) (string, error) {

	terrainString, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}
	return string(terrainString), nil
}

func (u *User) JsonToArray(arr string) ([]uint, error) {

	var array []uint
	err := json.Unmarshal([]byte(arr), &array)
	if err != nil {
		return array, err
	}
	return array, nil
}
