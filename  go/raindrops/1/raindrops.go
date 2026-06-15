package raindrops

import "fmt"

func Convert(number int) string {
	var raindrops = ""
	if number%3 == 0 {
		raindrops = "Pling"
	}
	if number%5 == 0 {
		raindrops += "Plang"
	}
	if number%7 == 0 {
		raindrops += "Plong"
	}
	if raindrops == "" {
		raindrops = fmt.Sprintf("%d", number)
	}
	return raindrops
}
