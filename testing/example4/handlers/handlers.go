package handlers

import (
	"encoding/json"
	"net/http"
)

//Routes set the route for web service
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

//SendJSON return simple json
func SendJSON(rw http.ResponseWriter, r *http.Request) {

	u := struct {
		Name  string
		Email string
	}{
		Name:  "andi",
		Email: "andi@andi",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(u)
}
