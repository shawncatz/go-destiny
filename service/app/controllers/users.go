package controllers

import (
	"github.com/revel/revel"
	"github.com/shawncatz/go-destiny/service/app/models"
	"github.com/shawncatz/go-destiny/service/app/routes"
)

type Users struct {
	App
}

func (c Users) Index() revel.Result {
	user := c.connected()

	return c.Render(user)
}

func (c Users) Login(email, password string, remember bool) revel.Result {
	revel.INFO.Printf("username=%s, password=%s, remember=%t", email, password, remember)

	user, err := models.UserFind(email)
	if err != nil {
		panic(err)
	}

	if ok := user.CheckPassword(password); !ok {
		c.Flash.Error("password incorrect")
		return c.Redirect(routes.App.Index())
	}

	c.Flash.Success("Welcome, " + user.Email)
	c.Session["user"] = user.Email

	return c.Redirect(routes.App.Index())
}

func (c Users) Register(email, password string) revel.Result {
	revel.INFO.Printf("username=%s, password=%s", email, password)

	user := &models.User{Email: email, Password: password}

	err := user.Save()
	if err != nil {
		panic(err)
	}

	c.Flash.Success("Welcome, " + user.Email)
	c.Session["user"] = user.Email

	return c.Redirect(routes.Users.Index())
}
