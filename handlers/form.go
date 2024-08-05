package handlers

import (
	"html/template"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	temp, err := template.ParseFiles("static/index.html")
	if err == nil {
		temp.Execute(w, nil)
	} else {
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}
