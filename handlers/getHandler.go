package handlers

import (
	"net/http"
	"text/template"
)

type PageData struct {
	Output []string
}

func HandleMainPage(w http.ResponseWriter, r *http.Request) {
	defer HandlePanic(w)
	data := PageData{
		Output: []string{}, // Initialize the output as empty
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		// Return a 404 status if the template file is not found or cannot be parsed
		http.Error(w, "Not Found: Template not found or invalid: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandlePanic(w http.ResponseWriter) {
	if r := recover(); r != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
