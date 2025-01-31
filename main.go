package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Suceess\n")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
}
