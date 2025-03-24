package middleware

import (
	"net/http"
	"AetherGo/internal/context"
	"log"
)

type MiddlewareFunc func(context.HandlerFunc) context.HandlerFunc

func Chain(middlewares ...MiddlewareFunc) context.HandlerFunc {
	return func(ctx *context.Context) {
		handler := func(ctx *context.Context) {} 
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		handler(ctx)
	}
}

func Logger(next context.HandlerFunc) context.HandlerFunc {
	return func(ctx *context.Context) {
		log.Printf("Request: %s %s", ctx.Request.Method, ctx.Request.URL.Path)
		next(ctx)
	}
}

func Recovery(next context.HandlerFunc) context.HandlerFunc {
	return func(ctx *context.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(ctx.Response, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		
	}
}
