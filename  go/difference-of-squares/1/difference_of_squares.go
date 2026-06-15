package differenceofsquares

import "math"

func SquareOfSum(n int) int {
	var sum = (float64(1+n) / float64(2)) * float64(n)
	return int(sum * sum)
}

func SumOfSquares(n int) int {
	var sum = 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	return int(math.Abs(float64(SumOfSquares(n) - SquareOfSum(n))))
}
