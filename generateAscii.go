package main

import (
	"fmt"
	"strings"
)

func generateAscii(TextInput string, loadBannerFile map[rune][]string) (string, error) {
	splitInput := strings.Split(TextInput, "\\n")
	fmt.Println(splitInput)

	var result strings.Builder

	for _, inputLine := range splitInput {
		if inputLine == "" {
			result.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for _, letter := range inputLine {
				// if letter < 32 || letter > 126 {
				// 	letter = '#'
				// }
				AsciiLines, ok := loadBannerFile[letter]
				if !ok {
					fmt.Println(fmt.Errorf("{letter} is not in the banner file and will be replaced by #"))
					letter = '#'
					AsciiLines, _ = loadBannerFile[letter]

				}
				result.WriteString(AsciiLines[i])

			}
			result.WriteString("\n")

		}

	}
	return result.String(), nil

}
