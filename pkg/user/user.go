package user

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
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

	key, err := strconv.Atoi(os.Getenv("HASH_KEY"))
	if err != nil {
		return "", err
	}
	fmt.Println(key)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), key)
	return string(bytes), err
}

func (u *User) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
