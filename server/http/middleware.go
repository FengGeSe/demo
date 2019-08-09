package http

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// interface
type HttpHandlerMiddleware func(http.Handler) http.Handler

// panic recover
func PanicRecoverMiddleware(h http.Handler) http.Handler {
	return panicHandler{h}
}

type panicHandler struct {
	next http.Handler
}

func (h panicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e, string(debug.Stack()))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			w.Write([]byte(`{"code":500,"msg":"Keep calm! Try again after prayer"}`))
		}
	}()
	h.next.ServeHTTP(w, r)
}
