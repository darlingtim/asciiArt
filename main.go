package main

import (
	"asciiArt/ascii"
	"asciiArt/handler"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

	// Web Entry Point
	fmt.Println("Starting Server...")

	router := http.NewServeMux()

	router.Handle("GET /", http.HandlerFunc(handler.HomeHandler))
	router.Handle("POST /ascii-art", http.HandlerFunc(handler.ProcessHandler))

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	startServerError := server.ListenAndServe()

	if startServerError != nil {
		fmt.Println(startServerError)
	} else {
		fmt.Printf("Server successfuly started on Port %v", &server.Addr)
	}

	//CLI Entry Point
	if len(os.Args) < 2 {
		fmt.Println(fmt.Errorf("USAGE: go run . Hello standard"))

	} else {
		TextInput := os.Args[1]
		BannerFile := "standard.txt"
		if len(os.Args) > 2 {
			BannerFile = os.Args[2]

		}
		if !(strings.HasSuffix(BannerFile, ".txt")) {
			BannerFile = BannerFile + ".txt"
		}

		loadBannerFile, err := ascii.LoadBanner(BannerFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		// split input Text input from command Line
		textInput := strings.ReplaceAll(TextInput, "\r\n", "\n")
		splitInput := strings.Split(textInput, "\\n")

		AsciiCharacter, err := ascii.GenerateAscii(splitInput, loadBannerFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(AsciiCharacter)
		return
	}
}
