package ascii

import (
	"fmt"
)

func validate(bannerData string) (bool, error) {
	if len(bannerData) == 0 {
		return false, fmt.Errorf("banner file is empty")
	}
	BannerLines := splitLines(bannerData)
	if len(BannerLines) != 855 {
		return false, fmt.Errorf("banner file is %d lines long instead of 855", len(BannerLines))
	}

	var invalidLines []int
	index := 0
	var widthLength int
	for i, lineWidth := range BannerLines {
		if i == index {
			index += 9
			widthLength = len(BannerLines[i+1])
			continue
		}
		if len(lineWidth) != widthLength {
			invalidLines = append(invalidLines, i+1)
			//fmt.Printf("%v + %d = %v\n", i, 1, len(lineWidth))
		}
	}
	if len(invalidLines) > 0 {
		return false, fmt.Errorf("the following lines do not have a consistent width with their counterpart character lines : %v", invalidLines)
	}

	return true, nil
}
