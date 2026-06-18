package isbnverifier

import "strings"

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func checkChar(c byte) int {
	if c == 'X' {
		return 10
	} else {
		return int(c - '0')
	}
}

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	sum := 0
	mul := 10
	for i := 0; i < 9; i++ {
		if isDigit(isbn[i]) {
			sum += int(isbn[i]-'0') * mul
			mul--
		} else { // invalid char
			return false
		}
	}
	sum += checkChar(isbn[9])
	return sum%11 == 0
}
