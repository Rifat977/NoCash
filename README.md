
# AetherGo

AetherGo is a lightweight, modular web framework written in Go, designed for building fast and scalable web applications.

## Architecture of AetherGo 🚀

AetherGo follows a lightweight, modular architecture inspired by MVC (Model-View-Controller) but optimized for Go. The framework is structured as follows:

### Model (M)

-   **Defined in:** `internal/db` and `nocash/models`
    
-   **Database ORM:** Uses GORM for Object-Relational Mapping (ORM).
    
-   **Functionality:** Simplifies database interactions, supporting migrations, relationships, and hooks.
    

### Application Core

-   **Defined in:** `internal/app`
    
-   **Functionality:** Handles the application lifecycle, bootstrapping the router, database, and middleware.
    

### Routes

-   **Defined in:** `nocash/routes/routes.go`
    
-   **Functionality:** Similar to Django’s `urls.py` or Laravel’s `routes/web.php`, registers HTTP endpoints and connects them to handlers.
    

### Handlers (Controller-Like)

-   **Defined in:** `nocash/handler/commonHandler.go`
    
-   **Functionality:** Handles incoming HTTP requests and maps them to appropriate views. Acts similarly to controllers in MVC.
    

### Views (V)

-   **Defined in:** `nocash/templates/`
    
-   **Functionality:** Renders HTML templates using `render.RenderTemplate`. Similar to Django’s MVT View layer or Laravel’s Blade templates.
    

### Middleware & Context Handling

-   **Defined in:** `internal/context` and `internal/middleware`
    
-   **Functionality:** Manages the request lifecycle, authentication, logging, and custom middleware. It allows for dependency injection, session management, and logging.

## Features

-   🚀 **High Performance:** Written in Go, compiled for low-latency and high throughput.
    
-   🛠️ **Built-in ORM (GORM):** Simplifies database interactions without needing raw SQL, supports migrations, relationships, and hooks.
    
-   🔄 **Custom Context System:** Provides middleware-like request handling, allowing dependency injection, session management, and logging.
    
-   🏗️ **Modular Design:** Encourages separating concerns between models, routes, and handlers, making the framework easy to extend.
    
-   🔥 **Live Reloading (Air):** Hot reloading during development, avoiding manual restarts when modifying code.
    
-   🛣️ **Simple Routing:** Easily map routes with parameters and custom handler functions.
    
-   🎨 **Template Rendering:** Renders HTML templates for dynamic content generation.

## Installation

```bash
go get github.com/rifat977/AetherGo

```

## Quick Start

### **1️⃣ Create a Basic Server**

```go
package main

import (
    "AetherGo/internal/app"
    "AetherGo/internal/context"
    "AetherGo/internal/render"
)

func main() {
    application := app.NewApp()

    // Register Routes
    RegisterRoutes(application)

    // Start the server
    application.Run()
}

func indexHandler(c *context.Context) {
    render.RenderJSON(c.Response, map[string]string{"message": "Hello, World!"})
}

func aboutHandler(c *context.Context) {
    name := c.Params["name"]
    render.RenderTemplate(c.Response, "nocash/templates/about.html", map[string]string{"Name": name})
}

```

### **2️⃣ Define Routes**

AetherGo provides a simple and flexible router with parameter support.

```go
package routes

import (
    "AetherGo/internal/app"
    "AetherGo/nocash/handler"
)

func RegisterRoutes(application *app.App) {
    application.Router.Add("GET", "/", handler.IndexHandler)
    application.Router.Add("GET", "/about/:name", handler.AboutHandler)
    application.Router.Add("GET", "/login", handler.LoginHandler)
}

```

### **3️⃣ Handler Implementation**

Handlers process HTTP requests using the `Context` object.

```go
package handler

import (
    "AetherGo/internal/context"
    "AetherGo/internal/render"
    "fmt"
)

func IndexHandler(c *context.Context) {
    render.RenderJSON(c.Response, map[string]string{"message": "Welcome to AetherGo!"})
}

func AboutHandler(c *context.Context) {
    name := c.Params["name"]
    render.RenderTemplate(c.Response, "nocash/templates/about.html", map[string]string{"Name": name})
}

func LoginHandler(c *context.Context) {
    fmt.Println("Login Handler triggered")
    if c.Request.Method == "GET" {
        render.RenderTemplate(c.Response, "nocash/templates/login.html", nil)
    }
}

```

## Middleware

AetherGo supports middleware for request processing, such as logging and panic recovery.

```go
app.Use(middleware.Logger)
app.Use(middleware.Recovery)

```

## Context

Each request handler receives a `Context` object that provides access to:

-   HTTP Request and Response
-   URL Parameters
-   Helper methods for JSON and HTML responses

```go
func userHandler(c *context.Context) {
    id := c.Params["id"]
    c.JSON(200, map[string]string{"id": id})
}

```

## Template Rendering

AetherGo supports HTML template rendering with data binding.

```go
render.RenderTemplate(c.Response, "templates/page.html", map[string]interface{}{
    "Title": "AetherGo Page",
    "User":  "John Doe",
})

```

## Database Integration

AetherGo includes built-in ORM support using **GORM** for seamless database interactions.

```go
import (
    "AetherGo/internal/db"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name  string
    Email string `gorm:"unique"`
}

func InitDB() {
    db := db.GetDB()
    db.AutoMigrate(&User{})
}

```

## Configuration

AetherGo provides an environment-based configuration system, allowing you to access settings dynamically.

```go
app.Config.GetPort()    // Returns configured port
app.Config.GetEnv()     // Returns current environment

```

## Future Features 🚀

🔹 **JWT Authentication** – Secure user authentication with JSON Web Tokens  
🔹 **WebSockets Support** – Enable real-time features like notifications & chat  
🔹 **Task Scheduling & Job Queue** – Background processing for better performance  
🔹 **GraphQL Support** – Alternative to REST for more flexible APIs  
🔹 **Rate Limiting** – Protect endpoints from abuse  
🔹 **CLI Tools** – Command-line utilities for managing the app easily

## Contributing

Contributions are welcome! Feel free to submit a Pull Request or raise an issue.

----------

### **License**

AetherGo is open-source and available under the MIT License.

Happy coding! 🚀
