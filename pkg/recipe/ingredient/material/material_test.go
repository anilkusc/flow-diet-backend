package material

import (
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Construct() (*gorm.DB, Material) {

	godotenv.Load("../../../.env")
	var material = Material{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Material_Name:       "bread",
		Material_Photo_Urls: []string{"S3URL1", "S3URL2"},
		Tags:                []string{"vegan"},
		Size:                200,
		Quantity:            "gram",
	}
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Material{})
	return db, material
}
func Destruct(db *gorm.DB) {
	db.Exec("DROP TABLE materials")
}
