package recipe

import (
	"errors"

	"gorm.io/gorm"
)

func (r *Recipe) Create(db *gorm.DB) error {
	var err error
	r.Ingredients_String, err = r.IngredientToJson(r.Ingredients)
	if err != nil {
		return err
	}
	r.Appropriate_Meals_String, err = r.ArrayToJson(r.Appropriate_Meals)
	if err != nil {
		return err
	}
	r.Photo_Urls_String, err = r.ArrayToJson(r.Photo_Urls)
	if err != nil {
		return err
	}
	r.Video_Urls_String, err = r.ArrayToJson(r.Video_Urls)
	if err != nil {
		return err
	}

	result := db.Create(r)
	return result.Error
}
func (r *Recipe) Read(db *gorm.DB) error {

	var err error
	var result *gorm.DB

	if r.ID != 0 {
		result = db.Where("id=?", r.ID).First(&r)
	} else {
		return errors.New("id cannot found")
	}
	r.Ingredients, err = r.JsonToIngredient(r.Ingredients_String)
	if err != nil {
		return err
	}
	r.Appropriate_Meals, err = r.JsonToArray(r.Appropriate_Meals_String)
	if err != nil {
		return err
	}
	r.Photo_Urls, err = r.JsonToArray(r.Photo_Urls_String)
	if err != nil {
		return err
	}
	r.Video_Urls, err = r.JsonToArray(r.Video_Urls_String)
	if err != nil {
		return err
	}

	return result.Error
}
func (r *Recipe) Update(db *gorm.DB) error {
	var err error
	r.Ingredients_String, err = r.IngredientToJson(r.Ingredients)
	if err != nil {
		return err
	}
	r.Appropriate_Meals_String, err = r.ArrayToJson(r.Appropriate_Meals)
	if err != nil {
		return err
	}
	r.Photo_Urls_String, err = r.ArrayToJson(r.Photo_Urls)
	if err != nil {
		return err
	}
	r.Photo_Urls_String, err = r.ArrayToJson(r.Photo_Urls)
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
		recipes[i].Ingredients, err = recipes[i].JsonToIngredient(recipes[i].Ingredients_String)
		if err != nil {
			return recipes, err
		}
		recipes[i].Appropriate_Meals, err = recipes[i].JsonToArray(recipes[i].Appropriate_Meals_String)
		if err != nil {
			return recipes, err
		}
		recipes[i].Photo_Urls, err = recipes[i].JsonToArray(recipes[i].Photo_Urls_String)
		if err != nil {
			return recipes, err
		}
		recipes[i].Video_Urls, err = recipes[i].JsonToArray(recipes[i].Video_Urls_String)
		if err != nil {
			return recipes, err
		}

	}
	return recipes, result.Error
}
