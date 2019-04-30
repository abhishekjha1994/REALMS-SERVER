package controllers

import (
	"fmt"

	"github.com/Realms-Server/app/models"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	App
}

func (a *Admin) Index() revel.Result {
	return a.Render()
}

// func (c Admin) checkUser() revel.Result {
// 	user := c.connected()
// 	if user == nil || user.UserType == "operator" {
// 		c.Flash.Error("Please log in first")
// 		return c.Redirect(App.ILoginPage)
// 	}

// 	return nil
// }

func (a Admin) Registration(username string, password string, email string, mobile_num string) revel.Result {

	if a.isUser(email) {
		return a.RenderJSON("User Already Exist")
		fmt.Println("user already there")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 16)

	status, _ := models.AddUser(username, email, hashed, mobile_num)
	if status {
		return a.RenderJSON("Success")
	}
	return a.RenderJSON("Failed")
	// log.Println("\n\ntry1\n\n")
}

func (a Admin) isUser(username string) bool {
	query := `select * from users where Username = '` + username + `'`
	users, err := a.Txn.Select(models.Users{}, query)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return false
	}
	return true
}
