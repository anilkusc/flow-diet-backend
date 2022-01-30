package user

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Construct() (*gorm.DB, User) {
	var user = User{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Username:         "testuser1",
		Name:             "test user",
		Password:         "testpass",
		Weight:           70,
		Height:           173,
		Age:              25,
		Diet:             "omnivor",
		Favorite_Recipes: []uint{1, 2, 3},
		Address:          "",
		Role:             "user",
	}
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&User{})
	return db, user
}
func Destruct() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Exec("DROP TABLE users")
}
