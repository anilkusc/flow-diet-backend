package shopping

import (
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Construct() (*gorm.DB, Shopping) {

	godotenv.Load("../../.env")
	var shopping = Shopping{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Ingredients: []int32{1, 2, 3},
		Start_Date:  1643743444,
		End_Date:    1643743448,
		User_Id:     1,
	}
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Shopping{})
	return db, shopping
}

func Destruct(db *gorm.DB) {
	db.Exec("DROP TABLE shoppings")
}
