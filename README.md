# AetherGo

AetherGo is a lightweight, modular web framework written in Go, designed for building fast and scalable web applications.

## Features

- 🚀 Lightweight and fast
- 🛣️ Simple routing with parameter support
- 🎨 Template rendering
- 🔄 Middleware support
- 📝 Built-in logging
- ⚙️ Environment-based configuration
- 🔒 Recovery middleware for panic handling
- 📦 JSON response support

## Installation

```bash
go get github.com/rifat977/AetherGo
```

## Quick Start

1. Create a basic server:

```go
package main

import (
    "AetherGo/internal/app"
    "AetherGo/internal/context"
    "AetherGo/internal/render"
)

func main() {
    app := app.NewApp()

    // Define routes
    app.Router.Add("GET", "/", indexHandler)
    app.Router.Add("GET", "/about/:name", aboutHandler)

    app.Run()
}

func indexHandler(c *context.Context) {
    render.RenderJSON(c.Response, map[string]string{"message": "Hello, World!"})
}

func aboutHandler(c *context.Context) {
    name := c.Params["name"]
    render.RenderTemplate(c.Response, "cmd/templates/about.html", map[string]string{"Name": name})
}
```

## Router

AetherGo provides a simple and flexible router with support for URL parameters:

```go
// Basic route
app.Router.Add("GET", "/", handler)

// Route with parameters
app.Router.Add("GET", "/users/:id", userHandler)
```

## Middleware

The framework supports middleware for request processing:

```go
app.Use(middleware.Logger)
app.Use(middleware.Recovery)
```

## Context

Each request handler receives a Context object that provides access to:
- HTTP Request and Response
- URL Parameters
- Helper methods for JSON and HTML responses

```go
func handler(c *context.Context) {
    // Access URL parameters
    id := c.Params["id"]
    
    // Send JSON response
    c.JSON(200, map[string]string{"id": id})
}
```

## Template Rendering

AetherGo supports HTML template rendering:

```go
render.RenderTemplate(c.Response, "templates/page.html", data)
```

## Configuration

Configuration is managed through environment variables and can be accessed through the Config struct:

```go
app.Config.GetPort()    // Returns configured port
app.Config.GetEnv()     // Returns current environment
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
