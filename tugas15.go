package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= 100; i++ {
		fmt.Fprintf(w, "%d ", i)
	}
}
func main() {

	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}
