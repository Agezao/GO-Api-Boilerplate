package routes

import mainRoute "./main"
import productsRoute "./products"
import "net/http"
import interceptor "../interceptors"
import . "../interceptors/trace"
import . "../interceptors/token"
import . "../interceptors/shaper"

func Bind() {
	intercept := interceptor.Chain(TraceInterceptor, ShaperInterceptor)
	interceptSecure := interceptor.Chain(intercept, TokenInterceptor)

	http.HandleFunc("/", intercept(mainRoute.Get))
	http.HandleFunc("/products", interceptSecure(productsRoute.Get))
}
