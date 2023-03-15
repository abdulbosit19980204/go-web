package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome my first page on golang uka ishlading ku %s", r.URL.Path)
}
func main() {
	http.HandleFunc("/uka", handlerFunc)
	http.ListenAndServe(":9090", nil)
}
