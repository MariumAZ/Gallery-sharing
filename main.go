package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func home(w http.ResponseWriter, r *http.Request) {
	// it is not obsvious that wwe are working always withh html
	// so we set the content header
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site")
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, "To get in touch please send an email please contact <a href=\"mail to : myriam.azzouz@holbertonschool.com\">myriam.azzouz@holbertonschool.com<a/>.")
}

func NotFound(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Sorry Page Requested Not Found")
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	r.HandleFunc("/", home)
    r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}