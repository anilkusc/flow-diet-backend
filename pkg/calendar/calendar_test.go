package calendar

import (
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Construct() (*gorm.DB, Calendar) {
	godotenv.Load("../../.env")
	var calendar = Calendar{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Recipe_Id:  1,
		User_Id:    1,
		Date_Epoch: 1643743444,
	}
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Calendar{})
	return db, calendar
}
func Destruct(db *gorm.DB) {
	db.Exec("DROP TABLE calendars")
}
