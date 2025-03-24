// package router

// import "net/http"

// type Router struct {
// 	routes map[string]http.HandlerFunc
// }

// func NewRouter() *Router {
// 	return &Router{routes: make(map[string]http.HandlerFunc)}
// }

package router

import (
	"net/http"
	"strings"

	"AetherGo/internal/context"
	"AetherGo/internal/log"
)

type route struct {
	method string
	pattern   string
	handler context.HandlerFunc
}

type Router struct {
	routes []route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Add(method, pattern string, handler context.HandlerFunc) {
	r.routes = append(r.routes, route{method, pattern, handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.method != req.Method {
			continue
		}

		params, ok := match(route.pattern, req.URL.Path)
		if ok {
			ctx := &context.Context {
				Response: w,
				Request: req,
				Params: params,
			}
			route.handler(ctx)
			log.Println("Route matched", route.method, route.pattern)
			return
		}
	}
	http.NotFound(w, req)
	log.Println("Route not found", req.Method, req.URL.Path)
}

func match(pattern, path string) (map[string]string, bool) {
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")
	if len(patternParts) != len(pathParts) {
		return nil, false
	}
	params := make(map[string]string)
	for i, part := range patternParts {
		if strings.HasPrefix(part, ":") {
			params[part[1:]] = pathParts[i]
		} else if part != pathParts[i] {
			return nil, false
		}
	}
	return params, true
}