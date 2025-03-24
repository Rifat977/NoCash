package handler

import (
	"AetherGo/internal/context"
	"AetherGo/internal/render"
)

// func IndexHandler(c *context.Context) {
// 	render.RenderJSON(c.Response, map[string]string{"message": "Hello, World!"})
// }

func LoginHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, "nocash/templates/login.html", nil)
	}
}

func AboutHandler(c *context.Context) {
	name := c.Params["name"]
	render.RenderTemplate(c.Response, "nocash/templates/about.html", map[string]string{"Name": name})
}

func RegisterHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, "nocash/templates/register.html", nil)
	}
}
