package productsRoute

import (
	"../../factories/response"
	. "../../models/product"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	product := Product{ID: 1, Name: "Test"}
	responseFactory.Ok(product, w)
}
