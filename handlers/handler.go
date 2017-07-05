package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// ServeResource serves static files(css and js) from server
func ServeResource(w http.ResponseWriter, r *http.Request) {
	path := "templates/" + r.URL.Path
	http.ServeFile(w, r, path)
}

// Home Serves the Landing page
func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index-slider.html")
	if err != nil {
		log.Println("Could not parse Home tpl", err)
	}
	t.Execute(w, nil)
}

// About Serves the about page
func About(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/about.html")
	if err != nil {
		log.Println("Could not parse About tpl", err)
	}
	t.Execute(w, nil)
}
