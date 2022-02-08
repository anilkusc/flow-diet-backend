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
		Measurement_Size:     200,
		Measurement_Quantity: "gram",
		Material_Photo_Urls:  []string{"S3URL1", "S3URL2"},
		Material_Name:        "Rice",
		IsExist:              false,
		IsOptional:           false,
	}
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Ingredient{})
	return db, ingredient
}
func Destruct(db *gorm.DB) {
	db.Exec("DROP TABLE ingredients")
}
