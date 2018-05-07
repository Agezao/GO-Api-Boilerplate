package productsRoute

import (
	"net/http"

	. "../../config"
	ctx "../../contexts/mongodb"
	"../../factories/response"
	. "../../models/product"
)

func Get(w http.ResponseWriter, r *http.Request) {
	ctx.Connect()
	products := ctx.Session.DB(Config.DbName).C("products")
	var result []Product
	products.Find(nil).All(&result)

	//product := Product{ID: 1, Name: "Test"}
	responseFactory.Ok(result, w)
}
