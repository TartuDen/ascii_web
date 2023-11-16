package main

import (
	"GO_ex3/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HandleMainPage)
	http.HandleFunc("/submit", handlers.HandleSubmitForm)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
