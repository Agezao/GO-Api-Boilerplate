package routes

import MainRoute "./main"
import ProductsRoute "./products"
import "net/http"
import interceptor "../interceptors"
import . "../interceptors/trace"
import . "../interceptors/token"
import . "../interceptors/shaper"

func Bind() {
	intercept := interceptor.Chain(TraceInterceptor, ShaperInterceptor)
	interceptSecure := interceptor.Chain(intercept, TokenInterceptor)

	http.HandleFunc("/", intercept(MainRoute.Get))
	http.HandleFunc("/products", interceptSecure(ProductsRoute.Get))
}
