package myappRoutes

import MainRoute "./main"
import ProductsRoute "./products"
import "net/http"

func Bind() {
	http.HandleFunc("/", MainRoute.Get)
	http.HandleFunc("/products", ProductsRoute.Get)
}
