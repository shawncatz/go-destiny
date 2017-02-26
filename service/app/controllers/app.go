package controllers

import (
	"github.com/revel/revel"
	"github.com/shawncatz/go-destiny/service/app/models"
	"github.com/shawncatz/go-destiny/service/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(routes.Users.Index())
	}
	return c.Render()
}

func (c App) Register() revel.Result {
	if c.connected() != nil {
		return c.Redirect(routes.Users.Index())
	}
	return c.Render()
}

func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}

func (c App) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}

	if email, ok := c.Session["user"]; ok {
		u, err := models.UserFind(email)
		if err != nil {
			return nil
		}

		c.RenderArgs["user"] = u
		return u
	}

	return nil
}
