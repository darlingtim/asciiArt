package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
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

		loadBannerFile, err := loadBanner(BannerFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		AsciiCharacter, err := generateAscii(TextInput, loadBannerFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(AsciiCharacter)
		return
	}
}
