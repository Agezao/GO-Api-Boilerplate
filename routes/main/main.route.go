package myappRoutesMain

import (
	"encoding/json"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Everything looks good aparently")
}
