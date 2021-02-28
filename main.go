package main

import (
	"encoding/json"
	api2 "github.com/sujanks/go-cql/src/api"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	api := api2.NewContoller()
	mux.HandleFunc("/api/query", func(writer http.ResponseWriter, r *http.Request) {
		response := api.Query(r.URL.Query().Get("q"))
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(response)
	})
	http.ListenAndServe(":8080", mux)
}


