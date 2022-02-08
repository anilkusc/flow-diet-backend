package shopping

import (
	"errors"

	"gorm.io/gorm"
)

func (s *Shopping) Create(db *gorm.DB) error {
	result := db.Create(s)
	return result.Error
}

func (s *Shopping) Read(db *gorm.DB) error {

	var result *gorm.DB

	if s.ID != 0 {
		result = db.Where("id=?", s.ID).First(&s)
	} else {
		return errors.New("id cannot found")
	}

	return result.Error
}

func (s *Shopping) Update(db *gorm.DB) error {
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

	var recipes []Shopping
	result := db.Where("user_id = ?", s.User_Id).Find(&recipes)

	return recipes, result.Error
}

func (s *Shopping) ListByDateInterval(db *gorm.DB) ([]Shopping, error) {

	var recipes []Shopping
	result := db.Where("user_id = ? AND start_date > ? AND end_date < ?", s.User_Id, s.Start_Date, s.End_Date).Find(&recipes)

	return recipes, result.Error
}
