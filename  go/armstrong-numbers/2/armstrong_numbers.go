package armstrongnumbers

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	s := strconv.Itoa(n)
	digits := len(s)
	sum := 0
	for i := 0; i < digits; i++ {
		sum += int(math.Pow(float64(s[i]-'0'), float64(digits)))
	}
	return n == sum
}
