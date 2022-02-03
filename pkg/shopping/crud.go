package shopping

import (
	"errors"

	"gorm.io/gorm"
)

func (s *Shopping) Create(db *gorm.DB) error {
	var err error
	s.Ingredients_String, err = s.ArrayToJson(s.Ingredients)
	if err != nil {
		return err
	}

	result := db.Create(s)
	return result.Error
}
func (s *Shopping) Read(db *gorm.DB) error {

	var err error
	var result *gorm.DB

	if s.ID != 0 {
		result = db.Where("id=?", s.ID).First(&s)
	} else {
		return errors.New("id cannot found")
	}
	//else {
	//result = db.Where("username=?", u.Username).First(&r)
	//}
	s.Ingredients, err = s.JsonToArray(s.Ingredients_String)
	if err != nil {
		return err
	}
	return result.Error
}
func (s *Shopping) Update(db *gorm.DB) error {
	var err error
	s.Ingredients_String, err = s.ArrayToJson(s.Ingredients)
	if err != nil {
		return err
	}
	result := db.Save(s)
	return result.Error
}
func (s *Shopping) Delete(db *gorm.DB) error {

	result := db.Delete(&Shopping{}, s.ID)
	return result.Error
}
func (s *Shopping) HardDelete(db *gorm.DB) error {

	result := db.Unscoped().Delete(&Shopping{}, s.ID)
	return result.Error
}

func (s *Shopping) List(db *gorm.DB) ([]Shopping, error) {
	var err error
	var recipes []Shopping
	result := db.Where("user_id = ?", s.User_Id).Find(&recipes)
	for i := range recipes {
		recipes[i].Ingredients, err = recipes[i].JsonToArray(recipes[i].Ingredients_String)
		if err != nil {
			return recipes, err
		}
	}
	return recipes, result.Error
}
