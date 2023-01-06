package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/hello-world", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := json.Marshal(map[string]string{"message": "hello world!"})
		w.Write(resp)
	}).Methods("GET")
	fmt.Println("listening on port 3000")
	http.ListenAndServe(":3000", r)
}
