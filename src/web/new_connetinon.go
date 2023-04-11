package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mihalic2040/Rosa-Router/src/manager"
)

func new_connection(w http.ResponseWriter, r *http.Request) {
	//p := manager.New_connection_out(Tunell: "localhost:9595")
	id := r.URL.Query().Get("id")
	fmt.Println(id)

	p := manager.Universal_out{Msg: "Ok..."}
	// Convert the person object to JSON
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(jsonBytes)
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the http.ResponseWriter
	w.Write(jsonBytes)
}
