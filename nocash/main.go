package main

import (
	"AetherGo/internal/app"
	"AetherGo/internal/db"

	"AetherGo/nocash/models"
	"AetherGo/nocash/routes"
)

func main() {
	app := app.NewApp()

	routes.RegisterRoutes(app) 
	
	db.AutoMigrate(&models.User{}) 

	app.Run()
}
