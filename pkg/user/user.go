package user

import (
	"errors"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        `json:"-" swaggerignore:"true"`
	Username          string         `gorm:"unique;not null" json:"username" example:"testuser"`
	Name              string         `json:"name" example:"test user"`
	Email             string         `gorm:"unique" json:"email" example:"test@test.com"`
	Phone             string         `gorm:"unique" json:"phone" example:"+905355353535"`
	Password          string         `json:"password,omitempty" example:"testpass"`
	Weight            uint8          `json:"weight" example:"70"`
	Height            uint8          `json:"height" example:"170"`
	Age               uint8          `json:"age" example:"25"`
	Gender            string         `json:"gender" example:"male"`  // male,female,other
	Diet_Level        uint           `json:"diet_level" example:"1"` //1: vegan ,2:vegaterian,3: omnivor ,4: carnivor
	Favorite_Recipes  pq.Int32Array  `json:"favorite_recipes" gorm:"type:int[]" example:"1,2,3"`
	Preferred_Meals   pq.StringArray `json:"preferred_meals" gorm:"type:text[]" example:"breakfast"`
	Likes             pq.StringArray `json:"likes" gorm:"type:text[]" example:"kebap,pizza"`
	Dislikes          pq.StringArray `json:"dislikes" gorm:"type:text[]" example:"onion"`
	Prohibits         pq.StringArray `json:"prohibits" gorm:"type:text[]" example:"sugar"`
	Address           string         `json:"address" example:"myadress 123121"`
	Role              string         `json:"role" example:"admin"` // root,admin,editor,user,anonymous
	Wants             string         `json:"wants" example:"gain"` // gain , lost , protect // (weights)
	Favorite_Cuisines pq.StringArray `json:"favorite_cousines" gorm:"type:text[]" example:"italian"`
	Blood_Group       string         `json:"blood_group" example:"A+"`   // A+,A- etc.
	Activity_Level    uint8          `json:"activity_level" example:"1"` //1,2,3,4,5
}

func (u *User) IsAuth(db *gorm.DB) (bool, error) {
	user := User{
		Username: u.Username,
	}

	err := user.Read(db)
	if err != nil {
		return false, err
	}
	if user.Username == u.Username {
		if u.CheckPasswordHash(u.Password, user.Password) {
			return true, nil
		}
	}
	return false, nil
}
func (u *User) Signup(db *gorm.DB) error {

	var err error
	//user := User{
	//	Username: u.Username,
	//}

	hashedPassword, err := u.HashPassword(u.Password)
	if err != nil {
		return err
	}
	nonHashedPassword := u.Password
	u.Password = hashedPassword
	err = u.Create(db)
	if err != nil {
		return err
	}
	u.Password = nonHashedPassword
	return nil

}
func (u *User) HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (u *User) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func (u *User) IsCredentialsExist(db *gorm.DB) error {
	var result int
	db.Raw("SELECT COUNT(username) FROM users WHERE username = ?;", u.Username).Scan(&result)
	if result != 0 {
		return errors.New("username is already exist")
	}
	db.Raw("SELECT COUNT(email) FROM users WHERE email = ?;", u.Email).Scan(&result)
	if result != 0 {
		return errors.New("email is already exist")
	}
	db.Raw("SELECT COUNT(phone) FROM users WHERE phone = ?;", u.Phone).Scan(&result)
	if result != 0 {
		return errors.New("phone number is already exist")
	}
	return nil
}
