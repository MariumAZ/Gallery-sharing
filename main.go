package main

import (
	"fmt"
	"net/http"
)


func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// it is not obsvious that wwe are working always withh html
	// so we set the content header
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site")
}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}