package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Congratulations, You have succesfully run your docker image.\n")
		fmt.Fprintf(w, "Please like, share and subscribe to this channel")
	})
	http.ListenAndServe(":8080", nil)
}
