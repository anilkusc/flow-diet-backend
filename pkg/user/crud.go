package user

import (
	"gorm.io/gorm"
)

func (u *User) Create(db *gorm.DB) error {
	var err error
	u.Favorite_Recipes_String, err = u.UintArrayToJson(u.Favorite_Recipes)
	if err != nil {
		return err
	}
	u.Preferred_Meals_String, err = u.ArrayToJson(u.Preferred_Meals)
	if err != nil {
		return err
	}
	u.Likes_String, err = u.ArrayToJson(u.Likes)
	if err != nil {
		return err
	}
	u.Dislikes_String, err = u.ArrayToJson(u.Dislikes)
	if err != nil {
		return err
	}
	u.Prohibits_String, err = u.ArrayToJson(u.Prohibits)
	if err != nil {
		return err
	}
	result := db.Create(u)
	return result.Error
}
func (u *User) Read(db *gorm.DB) error {
	var err error
	var result *gorm.DB
	if u.ID != 0 {
		result = db.Where("id=?", u.ID).First(&u)
	} else {
		result = db.Where("username=?", u.Username).First(&u)
	}

	u.Favorite_Recipes, err = u.JsonToUintArray(u.Favorite_Recipes_String)
	if err != nil {
		return err
	}
	u.Preferred_Meals, err = u.JsonToArray(u.Preferred_Meals_String)
	if err != nil {
		return err
	}
	u.Likes, err = u.JsonToArray(u.Likes_String)
	if err != nil {
		return err
	}
	u.Dislikes, err = u.JsonToArray(u.Dislikes_String)
	if err != nil {
		return err
	}
	u.Prohibits, err = u.JsonToArray(u.Prohibits_String)
	if err != nil {
		return err
	}

	return result.Error
}
func (u *User) Update(db *gorm.DB) error {
	var err error
	u.Favorite_Recipes_String, err = u.UintArrayToJson(u.Favorite_Recipes)
	if err != nil {
		return err
	}
	u.Preferred_Meals_String, err = u.ArrayToJson(u.Preferred_Meals)
	if err != nil {
		return err
	}
	u.Likes_String, err = u.ArrayToJson(u.Likes)
	if err != nil {
		return err
	}
	u.Dislikes_String, err = u.ArrayToJson(u.Dislikes)
	if err != nil {
		return err
	}
	u.Prohibits_String, err = u.ArrayToJson(u.Prohibits)
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

	result := db.Find(&users)
	for i := range users {
		users[i].Favorite_Recipes, err = users[i].JsonToUintArray(users[i].Favorite_Recipes_String)
		if err != nil {
			return users, err
		}
		users[i].Preferred_Meals, err = users[i].JsonToArray(users[i].Preferred_Meals_String)
		if err != nil {
			return users, err
		}
		users[i].Likes, err = users[i].JsonToArray(users[i].Likes_String)
		if err != nil {
			return users, err
		}
		users[i].Dislikes, err = users[i].JsonToArray(users[i].Dislikes_String)
		if err != nil {
			return users, err
		}
		users[i].Prohibits, err = users[i].JsonToArray(users[i].Prohibits_String)
		if err != nil {
			return users, err
		}
	}
	return users, result.Error
}
