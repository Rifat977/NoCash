package main

import (
	"AetherGo/internal/context"
	"AetherGo/internal/render"
)

func indexHandler(c *context.Context) {
	render.RenderJSON(c.Response, map[string]string{"message": "Hello, World!"})
}

func aboutHandler(c *context.Context) {
	name := c.Params["name"]
	render.RenderTemplate(c.Response, "cmd/templates/about.html", map[string]string{"Name": name})
}
