package controllers

import (
	"fmt"

	"github.com/revel/revel"
	"github.com/Realms-Server/app/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/Realms-Server/app/routes"
	//"github.com/Realms-Server/app/views"
)

type Reg struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_num"`
}
type login struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type App struct {
	GorpController
}

func (c App) Registration() revel.Result {
	user := c.connected()
	if user == nil {
		return c.Render()
	}
return c.Redirect(App.login)
}

func (c App) connected() *models.Users {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.Users)
	}
	if username, ok := c.Session["user"]; ok {
		return c.GetUser(username)
	}
	return nil
}
func (c App) App(user *models.Users, password string) bool {
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	if err == nil {
		return true
	}
	return false
}
func (c App) login() revel.Result {
	//add more code
	return c.Render()
}
func (c App)Homepage() revel.Result {
	//add more code
	return c.Render()
}

func (c App)Login(email string, password string, remember bool)revel.Result {
	user := c.GetUser(email)
	remember = true

	// if usertype == "" || usertype != "admin" || usertype != "operator" {
	// 	c.Flash.Error("Invalid user type: " + usertype)
	// }
	if user != nil {
		if checkPswd := c.App(user, password); checkPswd {
			c.Session["user"] = email
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			revel.INFO.Println("email IS:", user.Email)

			return c.Redirect(App.Homepage)
		}	

	}
	c.Flash.Out["email"] = email
	c.Flash.Error("Invalid email or password")
	return c.Redirect(App.login)
	//		c.Flash.Out["username"] = username
	//		c.Flash.Error("Invalid username or password")
	//		return c.Redirect(routes.App.LoginPage())
}

func (c App) GetUser(email string) *models.Users {
	query := `select * from users where email = '` + email + `'`
	var user models.Users
	err := models.Dbm.SelectOne(&user, query)
	if err != nil {
		fmt.Println("Error selecting user from DB for email : ", email, " Error: ", err)
		revel.ERROR.Println("Error selecting user from DB for email : ", email, " Error: ", err)
		return nil
	}
	return &user
}

func (c App) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.ViewArgs["user"] = user
	}
	return nil
}
//function for forget password
