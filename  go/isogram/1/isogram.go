package isogram

import "unicode"

func IsIsogram(word string) bool {
	foundChars := map[rune]bool{}
	for _, r := range word {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r) // unify letter case
			if foundChars[r] {
				return false
			}
			foundChars[r] = true
		}
	}
	return true
}
