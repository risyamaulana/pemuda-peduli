package web

import (
	"context"
	"fmt"
	"log"
	"pemuda-peduli/src/common/middleware"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// FastHTTPServer ...
type fastHTTPServer struct {
	server *fasthttp.Server
	router *router.Router
	port   int
}

func newFastHTTPServer(router *router.Router, port int) fastHTTPServer {
	f := fastHTTPServer{router: router, port: port}
	f.server = &fasthttp.Server{
		Handler:              middleware.Cors(fasthttp.CompressHandler(f.router.Handler)),
		ReadTimeout:          5 * time.Second,
		WriteTimeout:         10 * time.Second,
		MaxConnsPerIP:        50000,
		MaxRequestsPerConn:   50000,
		MaxKeepaliveDuration: 5 * time.Second,
	}

	return f
}

// Shutdown ...
func (f fastHTTPServer) Shutdown(ctx context.Context) {
	f.server.Shutdown()
}

// Listen ...
// Do not use *FastHTTPServer
func (f fastHTTPServer) Listen() {
	log.Printf("Web server started on localhost:%v\n", f.port)
	f.server.ListenAndServe(fmt.Sprintf(":%v", f.port))
}
