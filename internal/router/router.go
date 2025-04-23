package router

import (
	"net/http"
	"strings"

	"AetherGo/internal/context"
	"AetherGo/internal/log"
	"time"
)

type route struct {
	method  string
	pattern string
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
	rw := &responseWriter{ResponseWriter: w}

	start := time.Now()

	for _, route := range r.routes {
		if route.method != req.Method {
			continue
		}

		params, ok := match(route.pattern, req.URL.Path)
		if ok {
			ctx := &context.Context{
				Response: rw,
				Request:  req,
				Params:   params,
			}

			route.handler(ctx)

			duration := time.Since(start)

			log.Infof("\"%s %s %s\" %d %d %dms",
				req.Method,
				req.URL.Path,
				req.Proto,
				rw.status,
				rw.size,
				duration.Milliseconds())

			return
		}
	}

	duration := time.Since(start)
	log.Infof("\"%s %s %s\" 404 0 %dms",
		req.Method,
		req.URL.Path,
		req.Proto,
		duration.Milliseconds())
	http.NotFound(rw, req)
}

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
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
