package handler

import (
	"AetherGo/internal/context"
	"AetherGo/internal/db"
	"AetherGo/internal/model"
	"AetherGo/internal/render"
	"AetherGo/testapp/models"
	"net/http"

	// "text/template"
	"crypto/rand"
	"encoding/hex"
)

// func IndexHandler(c *context.Context) {
// 	render.RenderJSON(c.Response, map[string]string{"message": "Hello, World!"})
// }

func generateSessionToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

// write a funcitoin for get and check session

func LoginHandler(c *context.Context) {
	sessionToken := getSession(c, "session_token")
	if sessionToken != "" {
		var session models.Session
		if err := db.GetDB().Where("token = ?", sessionToken).First(&session).Error; err == nil {
			http.Redirect(c.Response, c.Request, "/dashboard", http.StatusSeeOther)
			return
		}
	}
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, nil, "testapp/templates/login.html")
	} else if c.Request.Method == "POST" {
		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")
		if email == "" || password == "" {
			http.Error(c.Response, "Email and password are required", http.StatusBadRequest)
			return
		}
		var user models.User
		if err := db.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
			http.Error(c.Response, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		checkPassword := model.CheckPasswordHash(password, user.Password)
		if !checkPassword {
			http.Error(c.Response, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		token, err := generateSessionToken()
		if err != nil {
			http.Error(c.Response, "Failed to generate session token", http.StatusInternalServerError)
			return
		}
		session := models.Session{
			UserID: user.ID,
			Token:  token,
			Expiry: 0,
		}
		if err := db.GetDB().Create(&session).Error; err != nil {
			http.Error(c.Response, "Failed to create session", http.StatusInternalServerError)
			return
		}
		http.SetCookie(c.Response, &http.Cookie{
			Name:  "session_token",
			Value: token,
			Path:  "/",
		})

		http.Redirect(c.Response, c.Request, "/dashboard", http.StatusSeeOther)

	}
}

func LogoutHandler(c *context.Context) {
	sessionToken := getSession(c, "session_token")
	if sessionToken != "" {
		var session models.Session
		if err := db.GetDB().Where("token = ?", sessionToken).First(&session).Error; err == nil {
			db.GetDB().Delete(&session)
		}
	}

	http.SetCookie(c.Response, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1,
	})

	http.Redirect(c.Response, c.Request, "/login", http.StatusSeeOther)
}

func RegisterHandler(c *context.Context) {
	sessionToken := getSession(c, "session_token")
	if sessionToken != "" {
		var session models.Session
		if err := db.GetDB().Where("token = ?", sessionToken).First(&session).Error; err == nil {
			http.Redirect(c.Response, c.Request, "/dashboard", http.StatusSeeOther)
			return
		}
	}

	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, nil, "testapp/templates/register.html")
		return
	} else if c.Request.Method == "POST" {
		if err := c.Request.ParseForm(); err != nil {
			render.RenderJSON(c.Response, map[string]string{"error": "Failed to parse form"})
			return
		}

		fullname := c.Request.FormValue("fullname")
		phone := c.Request.FormValue("phone")
		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")
		confirm := c.Request.FormValue("confirm_password")

		if fullname == "" || phone == "" || email == "" || password == "" || confirm == "" {
			http.Error(c.Response, "All fields are required", http.StatusBadRequest)
			return
		}

		if password != confirm {
			http.Error(c.Response, "Passwords do not match", http.StatusBadRequest)
			return
		}

		if len(password) < 6 {
			http.Error(c.Response, "Password must be at least 6 characters long", http.StatusBadRequest)
			return
		}

		encryptedPassword, err := model.EncryptPassword(password)
		if err != nil {
			http.Error(c.Response, "Failed to encrypt password", http.StatusInternalServerError)
			return
		}

		user := models.User{
			Name:     fullname,
			Phone:    phone,
			Email:    email,
			Password: encryptedPassword,
			Balance:  0,
		}

		if err := db.GetDB().Create(&user).Error; err != nil {
			http.Error(c.Response, "Failed to create user", http.StatusInternalServerError)
			return
		} else {
			token, err := generateSessionToken()
			if err != nil {
				http.Error(c.Response, "Failed to generate session token", http.StatusInternalServerError)
				return
			}

			session := models.Session{
				UserID: user.ID,
				Token:  token,
				Expiry: 0,
			}

			if err := db.GetDB().Create(&session).Error; err != nil {
				http.Error(c.Response, "Failed to create session", http.StatusInternalServerError)
				return
			}

			http.SetCookie(c.Response, &http.Cookie{
				Name:  "session_token",
				Value: token,
				Path:  "/",
			})

			http.Redirect(c.Response, c.Request, "/dashboard", http.StatusSeeOther)
		}

	}
}

func DashboardHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		data := map[string]interface{}{
			"currentPath": c.Request.URL.Path,
		}
		render.RenderTemplate(c.Response, data,
			"testapp/templates/partial/base.html",
			"testapp/templates/dashboard.html",
		)
	}
}

func CashInHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		data := map[string]interface{}{
			"currentPath": c.Request.URL.Path,
		}
		render.RenderTemplate(c.Response, data,
			"testapp/templates/partial/base.html",
			"testapp/templates/cash-in.html",
		)
	}
}

func CashOutHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		data := map[string]interface{}{
			"currentPath": c.Request.URL.Path,
		}
		render.RenderTemplate(c.Response, data,
			"testapp/templates/partial/base.html",
			"testapp/templates/cash-out.html",
		)
	}
}

func SendMoneyHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		data := map[string]interface{}{
			"currentPath": c.Request.URL.Path,
		}
		render.RenderTemplate(c.Response, data,
			"testapp/templates/partial/base.html",
			"testapp/templates/send-money.html",
		)
	}
}

func TransactionHistoryHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		data := map[string]interface{}{
			"currentPath": c.Request.URL.Path,
		}
		render.RenderTemplate(c.Response, data,
			"testapp/templates/partial/base.html",
			"testapp/templates/transactions.html",
		)
	}
}

func SettingsHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		data := map[string]interface{}{
			"currentPath": c.Request.URL.Path,
		}
		render.RenderTemplate(c.Response, data,
			"testapp/templates/partial/base.html",
			"testapp/templates/settings.html",
		)
	}
}
