package context

import (
	"net/http"
	"encoding/json"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params   map[string]string
}

type HandlerFunc func(*Context)

// show all logs
func (c *Context) JSON(status int, data interface{}) error {
	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.WriteHeader(status)
	return json.NewEncoder(c.Response).Encode(data)
}

func (c *Context) HTML(status int, html string) {
	c.Response.Header().Set("Content-Type", "text/html")
	c.Response.WriteHeader(status)
	c.Response.Write([]byte(html))
}
