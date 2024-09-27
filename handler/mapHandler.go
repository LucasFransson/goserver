package handler

import (
	"html/template"
	"log"
	"net/http"
)

func MapHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Map handler called")
	tmpl, err := template.ParseFiles("map.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
