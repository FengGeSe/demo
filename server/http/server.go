package http

import (
	"net"
	"net/http"
)

var mux = http.NewServeMux()

var httpServer = http.Server{Handler: mux}

func RegisterHandler(pattern string, h http.Handler) {
	mux.Handle(pattern, h)
}

func WrapMiddleware(handlerMap map[string]http.Handler, middlewares ...HttpHandlerMiddleware) map[string]http.Handler {
	m := map[string]http.Handler{}
	for k, h := range handlerMap {
		var newHandler http.Handler
		for _, m := range middlewares {
			newHandler = m(h)
		}
		m[k] = newHandler
	}
	return m
}

// http run
func Run(addr string, errc chan error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}
	errc <- httpServer.Serve(lis)
}
