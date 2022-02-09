package material

import (
	"gorm.io/gorm"
)

func (m *Material) Create(db *gorm.DB) error {
	result := db.Create(m)
	return result.Error
}
func (m *Material) Read(db *gorm.DB) error {

	result := db.Where("id=?", m.ID).First(&m)
	return result.Error
}
func (m *Material) Update(db *gorm.DB) error {

	result := db.Save(m)
	return result.Error
}
func (m *Material) Delete(db *gorm.DB) error {

	result := db.Delete(&Material{}, m.ID)
	return result.Error
}
func (m *Material) HardDelete(db *gorm.DB) error {

	result := db.Unscoped().Delete(&Material{}, m.ID)
	return result.Error
}

func (m *Material) List(db *gorm.DB) ([]Material, error) {

	var materials []Material
	result := db.Find(&materials)

	return materials, result.Error
}
