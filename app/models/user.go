package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/revel/revel"
)

type Users struct {
	Username       string `db:"username"`
	Password       string `db:"password"`
	HashedPassword []byte
	Email          string `db:"email"`
	MobileNumber   string `db:"mobile_num"`
}

func (u *Users) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *Users) Validate(v *revel.Validation) {
	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).
		Key("user.Password")

	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}

//Createsa a new user in DB
func AddUser(username string, email string, hashed []byte, mobile_num string) (bool, error) {
	userDetails := Users{Username: username, Email: email, HashedPassword: hashed, mobile_num: mobile_num}
	err := Dbm.Insert(&userDetails)
	if err != nil {
		log.Println("Error creating New User : ", err)
		return false, err
	}
	return true, nil

}
