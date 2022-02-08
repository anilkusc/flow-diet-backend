package recipe

import (
	"errors"

	"gorm.io/gorm"
)

func (r *Recipe) Create(db *gorm.DB) error {
	result := db.Create(r)
	return result.Error
}
func (r *Recipe) Read(db *gorm.DB) error {

	var result *gorm.DB

	if r.ID != 0 {
		result = db.Where("id=?", r.ID).First(&r)
	} else {
		return errors.New("id cannot found")
	}
	return result.Error
}
func (r *Recipe) Update(db *gorm.DB) error {

	result := db.Save(r)
	return result.Error
}
func (r *Recipe) Delete(db *gorm.DB) error {

	result := db.Delete(&Recipe{}, r.ID)
	return result.Error
}
func (r *Recipe) HardDelete(db *gorm.DB) error {

	result := db.Unscoped().Delete(&Recipe{}, r.ID)
	return result.Error
}

func (r *Recipe) List(db *gorm.DB) ([]Recipe, error) {

	var recipes []Recipe
	result := db.Find(&recipes)

	return recipes, result.Error
}

func (r *Recipe) ListWithLimit(db *gorm.DB, limit int) ([]Recipe, error) {

	var recipes []Recipe
	result := db.Limit(limit).Find(&recipes)

	return recipes, result.Error
}
