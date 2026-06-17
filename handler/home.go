package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	view, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, fmt.Sprint(fmt.Errorf("%v", err)), http.StatusInternalServerError)
		return
	}

	view.Execute(w, nil)
}
