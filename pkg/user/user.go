package user

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model              `json:"-" swaggerignore:"true"`
	Username                string `gorm:"unique;not null" json:"username" example:"testuser"`
	Name                    string `json:"name" example:"test user"`
	Password                string `json:"password" example:"testpass"`
	Weight                  uint8  `json:"weight" example:"70"`
	Height                  uint8  `json:"height" example:"170"`
	Age                     uint8  `json:"age" example:"25"`
	Diet                    string `json:"diet" example:"omnivor"` //vegaterian , vegan , omnivor , carnivor
	Favorite_Recipes        []uint `json:"favorite_recipes" gorm:"-"`
	Favorite_Recipes_String string `json:"-" swaggerignore:"true"`
	Address                 string `json:"address" example:"myadress 123121"`
	Role                    string `json:"role" example:"user"` // user,admin,dietician,editor
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

func (u *User) IsAuth(db *gorm.DB) (bool, error) {
	user := User{
		Username: u.Username,
	}

	err := user.Read(db)
	if err != nil {
		return false, err
	}
	if user.Username == u.Username {
		if u.CheckPasswordHash(u.Password, user.Password) {
			return true, nil
		}
	}
	return false, nil
}
func (u *User) Signup(db *gorm.DB) error {
	var err error
	hashedPassword, err := u.HashPassword(u.Password)
	if err != nil {
		return err
	}
	nonHashedPassword := u.Password
	u.Password = hashedPassword
	err = u.Create(db)
	if err != nil {
		return err
	}
	u.Password = nonHashedPassword
	return nil

}
func (u *User) HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (u *User) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
