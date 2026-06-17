package ascii

import (
	"strings"
)

func splitLines(bannerData string) []string {
	return strings.Split(bannerData, "\n")

}
