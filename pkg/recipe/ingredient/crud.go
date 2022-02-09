package ingredient

import (
	"gorm.io/gorm"
)

func (i *Ingredient) Create(db *gorm.DB) error {
	result := db.Create(i)
	return result.Error
}
func (i *Ingredient) Read(db *gorm.DB) error {

	result := db.Where("id=?", i.ID).First(&i)
	return result.Error
}
func (i *Ingredient) Update(db *gorm.DB) error {

	result := db.Save(i)
	return result.Error
}
func (i *Ingredient) Delete(db *gorm.DB) error {

	result := db.Delete(&Ingredient{}, i.ID)
	return result.Error
}
func (i *Ingredient) HardDelete(db *gorm.DB) error {

	result := db.Unscoped().Delete(&Ingredient{}, i.ID)
	return result.Error
}

func (i *Ingredient) List(db *gorm.DB) ([]Ingredient, error) {

	var Ingredients []Ingredient
	result := db.Find(&Ingredients)

	return Ingredients, result.Error
}

func (i *Ingredient) ListWithLimit(db *gorm.DB, limit int) ([]Ingredient, error) {

	var Ingredients []Ingredient
	result := db.Limit(limit).Find(&Ingredients)

	return Ingredients, result.Error
}
