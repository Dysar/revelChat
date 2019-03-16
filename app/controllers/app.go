package controllers

import (
	"github.com/revel/revel"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {

	return c.Render()
}

func (c Application) Destroy() {
	c.Controller.Destroy()
}

func (c Application) EnterDemo(user, phone string) revel.Result {

	c.Validation.Required(user)
	c.Validation.Required(phone)

	if c.Validation.HasErrors() {
		c.Flash.Error("Please choose a nick name.")
		return c.Redirect(Application.Index)
	}

	//switch demo {
	// case "refresh":
	// 	return c.Redirect("/refresh?user=%s", user)
	// case "longpolling":
	// 	return c.Redirect("/longpolling/room?user=%s", user)
	// case "websocket":
	return c.Redirect("/websocket/room?user=%s", user)
	//}
}
