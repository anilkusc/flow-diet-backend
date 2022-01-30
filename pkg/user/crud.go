package user

import (
	"gorm.io/gorm"
)

func (u *User) Create(db *gorm.DB) error {
	var err error
	u.Favorite_Recipes_String, err = u.ArrayToJson(u.Favorite_Recipes)
	if err != nil {
		return err
	}
	result := db.Create(u)
	return result.Error
}
func (u *User) Read(db *gorm.DB) error {
	var err error
	result := db.First(&u)
	u.Favorite_Recipes, err = u.JsonToArray(u.Favorite_Recipes_String)
	if err != nil {
		return err
	}

	return result.Error
}
func (u *User) Update(db *gorm.DB) error {
	var err error
	u.Favorite_Recipes_String, err = u.ArrayToJson(u.Favorite_Recipes)
	if err != nil {
		return err
	}
	result := db.Save(u)
	return result.Error
}
func (u *User) Delete(db *gorm.DB) error {

	result := db.Delete(&User{}, u.ID)
	return result.Error
}
func (u *User) HardDelete(db *gorm.DB) error {

	result := db.Unscoped().Delete(&User{}, u.ID)
	return result.Error
}
func (u *User) List(db *gorm.DB) ([]User, error) {
	var err error
	var users []User

	//result := db.Where("game_id = ? ", u.ID).Find(&users)
	result := db.Find(&users)
	for i := range users {
		users[i].Favorite_Recipes, err = users[i].JsonToArray(users[i].Favorite_Recipes_String)
		if err != nil {
			return users, err
		}
	}
	return users, result.Error
}
