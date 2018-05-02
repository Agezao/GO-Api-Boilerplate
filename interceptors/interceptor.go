package interceptor

import (
	"net/http"
)

type Interceptor func(http.HandlerFunc) http.HandlerFunc

func Chain(interceptors ...Interceptor) Interceptor {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(interceptors) - 1; i >= 0; i-- {
				last = interceptors[i](last)
			}

			last(w, r)
		}
	}
}
