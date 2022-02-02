package recipe

import "gorm.io/gorm"

func (r *Recipe) Create(db *gorm.DB) error {
	var err error
	r.Ingredients_String, err = r.ArrayToJson(r.Ingredients)
	if err != nil {
		return err
	}

	result := db.Create(r)
	return result.Error
}
func (r *Recipe) Read(db *gorm.DB) error {

	var err error
	r.Ingredients, err = r.JsonToArray(r.Ingredients_String)
	if err != nil {
		return err
	}

	result := db.First(&r)
	return result.Error
}
func (r *Recipe) Update(db *gorm.DB) error {
	var err error
	r.Ingredients_String, err = r.ArrayToJson(r.Ingredients)
	if err != nil {
		return err
	}
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
	var err error
	var recipes []Recipe
	result := db.Find(&recipes)
	for i := range recipes {
		recipes[i].Ingredients, err = recipes[i].JsonToArray(recipes[i].Ingredients_String)
		if err != nil {
			return recipes, err
		}
	}
	return recipes, result.Error
}
