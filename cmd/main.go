package main

import (
	"AetherGo/internal/app"
	"AetherGo/internal/db"
)

func main() {
	app := app.NewApp()

	RegisterRoutes(app)

	db.AutoMigrate(&User{})

	app.Run()
}
