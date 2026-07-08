package pangram

import "unicode"

func IsPangram(input string) bool {
	chars := make(map[rune]int, 0)
	for _, c := range input {
		if unicode.IsLetter(c) {
			c = unicode.ToLower(c)
			if chars[c] == 0 {
				chars[c]++
			}
		}
	}
	return len(chars) == 26
}
