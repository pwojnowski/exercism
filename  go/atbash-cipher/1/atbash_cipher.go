package atbashcipher

import "unicode"

func Atbash(s string) string {
	encoded := make([]rune, 0)
	groupCount := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			encoded = append(encoded, r)
			groupCount++
		} else if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			offset := (r - 'a')
			ciphered := ('z' - offset)
			encoded = append(encoded, ciphered)
			groupCount++
		}
		if groupCount == 5 {
			encoded = append(encoded, ' ')
			groupCount = 0
		}
	}
	last := len(encoded) - 1
	if encoded[last] == ' ' {
		encoded = encoded[:last]
	}
	return string(encoded)
}
