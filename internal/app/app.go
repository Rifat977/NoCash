package app

import (
	"net/http"
	"sync"

	"AetherGo/internal/config"
	"AetherGo/internal/log"
	"AetherGo/internal/middleware"
	"AetherGo/internal/router"

	"fmt"
	"os"
	"path/filepath"
)

type App struct {
	Config      *config.Config
	Router      *router.Router
	middlewares []middleware.MiddlewareFunc
	mu          sync.RWMutex
}

func NewApp(cfg *config.Config) *App {
	app := &App{
		Config: cfg,
		Router: router.NewRouter(),
	}
	return app
}

func (a *App) Use(mw middleware.MiddlewareFunc) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.middlewares = append(a.middlewares, mw)
}

func (a *App) Run() error {
	port := a.Config.GetPort()

	if a.Config.GetEnv() == "development" {
		log.Infof("Starting development server at http://127.0.0.1%s", port)
	} else {
		log.Infof("Starting production server at http://127.0.0.1%s", port)
	}

	log.Infof("Quit the server with CONTROL-C.\n\n")

	return http.ListenAndServe(port, a.Router)
}

func CreateNewProject(projectName string) error {
	dirs := []string{
		projectName,
		filepath.Join(projectName, "routes"),
		filepath.Join(projectName, "handler"),
		filepath.Join(projectName, "models"),
		filepath.Join(projectName, "templates"),
	}

	// Create directories
	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	// Create a sample `main.go` file
	mainFilePath := filepath.Join(projectName, "main.go")
	mainFileContent := `package main

import (
	"` + projectName + `/internal/app"
	"` + projectName + `/routes"
)

func main() {
	app := app.NewApp()
	routes.RegisterRoutes(app)
	app.Run()
}`

	err := os.WriteFile(mainFilePath, []byte(mainFileContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to create main.go: %v", err)
	}

	// Create a sample `commonHandler.go` file
	handlerFilePath := filepath.Join(projectName, "handler", "commonHandler.go")
	handlerFileContent := `package handler

import (
	"` + projectName + `/internal/context"
	"` + projectName + `/internal/render"
)

func IndexHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, "templates/index.html", nil)
	}
}`

	err = os.WriteFile(handlerFilePath, []byte(handlerFileContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to create handler/commonHandler.go: %v", err)
	}

	// Create a sample `User.go` model file
	modelFilePath := filepath.Join(projectName, "models", "User.go")
	modelFileContent := `package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string ` + "`gorm:\"unique\"`" + `
}`

	err = os.WriteFile(modelFilePath, []byte(modelFileContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to create models/User.go: %v", err)
	}

	// Create a sample `routes.go` file
	routesFilePath := filepath.Join(projectName, "routes", "routes.go")
	routesFileContent := `package routes

import (
	"` + projectName + `/internal/app"
	"` + projectName + `/handler"
)

func RegisterRoutes(application *app.App) {
	application.Router.Add("GET", "/", handler.IndexHandler)
}`

	err = os.WriteFile(routesFilePath, []byte(routesFileContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to create routes/routes.go: %v", err)
	}

	fmt.Printf("Project '%s' created successfully!\n", projectName)
	return nil
}
