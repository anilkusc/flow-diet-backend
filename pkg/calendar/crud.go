package calendar

import "gorm.io/gorm"

func (c *Calendar) Create(db *gorm.DB) error {
	result := db.Create(c)
	return result.Error
}
func (c *Calendar) Read(db *gorm.DB) error {
	result := db.First(&c)

	return result.Error
}
func (c *Calendar) Update(db *gorm.DB) error {
	result := db.Save(c)
	return result.Error
}
func (c *Calendar) Delete(db *gorm.DB) error {

	result := db.Delete(&Calendar{}, c.ID)
	return result.Error
}
func (c *Calendar) HardDelete(db *gorm.DB) error {

	result := db.Unscoped().Delete(&Calendar{}, c.ID)
	return result.Error
}
func (c *Calendar) List(db *gorm.DB) ([]Calendar, error) {

	var calendars []Calendar
	result := db.Where("user_id = ? ", c.User_Id).Find(&calendars)
	return calendars, result.Error
}
