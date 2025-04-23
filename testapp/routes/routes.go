package routes

import (
	"AetherGo/internal/app"

	"AetherGo/testapp/handler"
)

func RegisterRoutes(application *app.App) {
	application.Router.Add("GET", "/", handler.LoginHandler)

	application.Router.Add("GET", "/register", handler.RegisterHandler)
	application.Router.Add("POST", "/register", handler.RegisterHandler)

	application.Router.Add("GET", "/logout", handler.LogoutHandler)

	application.Router.Add("GET", "/dashboard", handler.DashboardHandler)
	application.Router.Add("GET", "/cash-in", handler.CashInHandler)
	application.Router.Add("GET", "/cash-out", handler.CashOutHandler)
	application.Router.Add("GET", "/send-money", handler.SendMoneyHandler)
	application.Router.Add("GET", "/transactions", handler.TransactionHistoryHandler)
	application.Router.Add("GET", "/settings", handler.SettingsHandler)

}
