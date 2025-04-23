package handler

import (
	"AetherGo/internal/context"
	"fmt"
)

// func getSession(ctx *context.Context, session_name string) string {
// 	cookie, err := ctx.Request.Cookie(session_name)
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			log.Debug("Session token not found in cookies")
// 		}
// 		log.Error("Error retrieving session token cookie:", err)
// 	}

// 	log.Debug("Session token found: ", cookie.Value)
// 	return cookie.Value
// }

func getSession(c *context.Context, key string) string {
	cookie, err := c.Request.Cookie(key)
	if err != nil {
		// Log and return empty string
		fmt.Printf("[ERROR] Error retrieving session token cookie: %v", err)
		return ""
	}
	return cookie.Value
}
