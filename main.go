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
var signView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signView.Render(w, nil))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Sorry Page Requested Not Found")
}

func main() {

	homeView    = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	signView    = views.NewView("bootstrap", "views/signup.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", signup)
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}