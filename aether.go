package AetherGo

import (
	"AetherGo/internal/app"
	"AetherGo/internal/config"
	"AetherGo/internal/db"
	"AetherGo/internal/log"
)

type App = app.App

func Bootstrap(appName string, models []interface{}, routeRegistrar func(*App), overrides ...config.Config) *App {
	cfg := configure(overrides...)
	app := app.NewApp(cfg)

	db.ConnectDB(appName)

	log.Infof("AetherGo version 1.0, using environment '%s'", appName)

	routeRegistrar(app)

	db.AutoMigrate(models...)

	return app
}

func configure(overrides ...config.Config) *config.Config {
	cfg := &config.Config{
		Port:         "8000",
		Env:          "development",
		TemplatesDir: "templates",
		StaticDir:    "static",
	}

	if len(overrides) > 0 {
		override := overrides[0]
		if override.Port != "" {
			cfg.Port = override.Port
		}
		if override.Env != "" {
			cfg.Env = override.Env
		}
		if override.TemplatesDir != "" {
			cfg.TemplatesDir = override.TemplatesDir
		}
		if override.StaticDir != "" {
			cfg.StaticDir = override.StaticDir
		}
	}

	return cfg
}

func RegisterRoutes(app *app.App, registerFunc func(*app.App)) {
	registerFunc(app)
}

func Run(app *app.App) {
	app.Run()
}
