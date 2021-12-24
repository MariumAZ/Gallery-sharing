package main

import (
	"Gallery-sharing/views"
	"fmt"
	_ "html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	// it is not obsvious that wwe are working always withh html
	// so we set the content header
	w.Header().Set("Content-Type", "text/html")
	homeTemplate := homeView.Template
	err := homeTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	contactTemplate := contactView.Template
	err := contactTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Sorry Page Requested Not Found")
}

func main() {
	// keep track of the template being parsed
	//var err error
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
