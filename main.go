package main

import (
	"encoding/json"
	"github.com/benjamint08/vite-react-go-template/helpers"
	"net/http"
)

func main() {
	helpers.CheckFlags()
	
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello from Go!"})
	})

	http.ListenAndServe(":8080", nil)
}
