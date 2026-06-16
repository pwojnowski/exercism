package luhn

import "strings"

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func getValue(digitChar byte, double bool) int {
	x := int(digitChar - '0')
	if double {
		x *= 2
		if x > 9 {
			x -= 9
		}
	}
	return x
}

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	double := false
	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		if !isDigit(id[i]) {
			return false
		}
		sum += getValue(id[i], double)
		double = !double
	}
	return sum%10 == 0
}
