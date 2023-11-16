package handlers

import (
	"GO_ex3/functions"
	"net/http"
	"strings"
	"text/template"
)

func HandleSubmitForm(w http.ResponseWriter, r *http.Request) {
	defer HandlePanic(w)
	if r.Method != http.MethodPost {
		// Return Bad Request status if the request method is incorrect
		http.Error(w, "Bad Request", http.StatusMethodNotAllowed)
		return
	}

	inputTextFromPage := r.FormValue("textInput")
	// if inputTextFromPage == "" {
	// 	// Return Bad Request status if the inputText is empty (as an example)
	// 	http.Error(w, "Bad Request: inputText is required", http.StatusBadRequest)
	// 	return
	// }
	banner := r.FormValue("banner")
	banner = banner + ".txt"

	// color := r.FormValue("textColor")
	// fmt.Println("color is:", color)
	inputTextFromPage = strings.ReplaceAll(inputTextFromPage, "\r", "")
	inputTextAfterASCII := functions.RunFunction(inputTextFromPage, banner, "")
	outputText := inputTextAfterASCII

	data := PageData{
		Output: outputText,
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		// Return a 404 status if the template file is not found or cannot be parsed
		http.Error(w, "Not Found: Template not found or invalid", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
