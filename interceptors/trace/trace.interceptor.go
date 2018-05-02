package traceInterceptor

import (
	"fmt"
	"net/http"
)

func TraceInterceptor(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("incoming ", r.Method, " at ", r.URL)
		next.ServeHTTP(w, r)
	}
}
