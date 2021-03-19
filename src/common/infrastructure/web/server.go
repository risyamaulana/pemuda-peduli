package web

import (
	"context"

	"github.com/fasthttp/router"
)

// Server ...
type Server interface {
	Listen()
	Shutdown(context.Context)
}

// NewWebServer ...
func NewWebServer(route *router.Router, port int) (Server, error) {
	return newFastHTTPServer(route, port), nil
}

// NewRouter ...
// [TODO] add alternative router implementation here
func NewRouter() *router.Router {
	router := router.New()
	return router
}
