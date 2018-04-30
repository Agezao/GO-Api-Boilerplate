package myappRoutesProducts

import (
	. "../../models/product"
	"encoding/json"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product := Product{ID: 1, Name: "Test"}
	json.NewEncoder(w).Encode(product)
}
