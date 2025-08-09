package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world", r.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
