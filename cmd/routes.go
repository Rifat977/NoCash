package main

import (
	"AetherGo/internal/app"
)

func RegisterRoutes(application *app.App) {
	application.Router.Add("GET", "/", indexHandler)
	application.Router.Add("GET", "/about/:name", aboutHandler)
}
