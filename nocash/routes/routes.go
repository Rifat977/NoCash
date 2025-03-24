package routes

import (
	"AetherGo/internal/app"

	"AetherGo/nocash/handler"
)

func RegisterRoutes(application *app.App) {
	application.Router.Add("GET", "/", handler.LoginHandler)
	application.Router.Add("GET", "/register", handler.RegisterHandler)
	application.Router.Add("GET", "/about/:name", handler.AboutHandler)
}
