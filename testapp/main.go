package main

import (
	"AetherGo"
	"AetherGo/testapp/models"
	"AetherGo/testapp/routes"
	"log"
)

func main() {
	app := AetherGo.Bootstrap("testapp", models.GetAll(), routes.RegisterRoutes)
	AetherGo.Run(app)
	log.Println("🌟 Application started!")
}
