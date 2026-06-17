package handler

import (
	"asciiArt/ascii"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type AsciiData struct {
	Text        string
	Banner      string
	AsciiResult string
}

var Data *AsciiData

func ProcessHandler(w http.ResponseWriter, r *http.Request) {

	view, err := template.ParseFiles("templates/home.html")
	// if err != nil {
	// 	http.Error(w, fmt.Sprint(fmt.Errorf("%v", err)), http.StatusInternalServerError)
	// 	return
	// }

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	input := strings.ReplaceAll(text, "\r\n", "\n")
	splitInput := strings.Split(input, "\n")
	fmt.Println(splitInput)

	loadedBanner, _ := ascii.LoadBanner(banner)
	asciiResult, err := ascii.GenerateAscii(splitInput, loadedBanner)
	if err != nil {
		asciiResult = fmt.Sprint(fmt.Errorf("%v", err))
	}

	Data = &AsciiData{
		Text:        text,
		Banner:      banner,
		AsciiResult: asciiResult,
	}
	view.Execute(w, Data)

}
