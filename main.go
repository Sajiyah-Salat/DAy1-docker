package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Saurav, Welcome to the first session of Docker.")
	})
	http.ListenAndServe(":8080", nil)
}
