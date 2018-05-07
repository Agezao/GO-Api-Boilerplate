package productsRoute

import (
	"net/http"

	"../../factories/response"
	productService "../../services/products"
)

func Get(w http.ResponseWriter, r *http.Request) {
	result := productService.Get(nil)

	responseFactory.Ok(result, w)
}
