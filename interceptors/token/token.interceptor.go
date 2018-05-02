package tokenInterceptor

import (
	"../../factories/response"
	"net/http"
)

func TokenInterceptor(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil || r.Header["Authorization"][0] != "1234" {
			responseFactory.Fail(-1, "Unauthorized", w)
			return
		}

		next.ServeHTTP(w, r)
	}
}
