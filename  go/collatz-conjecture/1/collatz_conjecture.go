package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var steps = 0
	if n <= 0 {
		return steps, errors.New("n must be positive")
	}
	for 1 < n {
		steps++
		if n&1 == 1 {
			n = (n * 3) + 1
		} else {
			n /= 2
		}
	}
	return steps, nil
}
