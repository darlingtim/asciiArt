package main

import (
	"os"
)

func loadBanner(bannerFile string) (map[rune][]string, error) {
	banner := make(map[rune][]string)
	bannerData, err := os.ReadFile(bannerFile)
	if err != nil {
		return nil, err
	}

	//  validate file content
	validate, validateError := validate(string(bannerData))
	if !validate {
		return nil, validateError
	}

	BannerLines := splitLines(string(bannerData))

	var characterLines []string
	character := ' '
	for _, line := range BannerLines {
		if line == "" && len(characterLines) == 0 {
			continue
		}
		characterLines = append(characterLines, line)
		if len(characterLines) == 8 {
			banner[character] = characterLines
			characterLines = nil
			character++
		}

	}

	return banner, nil
}
