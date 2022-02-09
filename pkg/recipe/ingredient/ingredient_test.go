package ingredient

import (
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Construct() (*gorm.DB, Ingredient) {

	godotenv.Load("../../../.env")
	var ingredient = Ingredient{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Material_Id: 1,
		IsOptional:  false,
	}
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Ingredient{})
	return db, ingredient
}
func Destruct(db *gorm.DB) {
	db.Exec("DROP TABLE ingredients")
}
