package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	if amount, err := fc.FodderAmount(cows); err != nil {
		return 0, err
	} else if factor, err := fc.FatteningFactor(); err != nil {
		return 0, err
	} else {
		return (amount * factor) / float64(cows), nil
	}
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fc, cows)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	Cows    int
	Message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.Cows, err.Message)
}

func ValidateNumberOfCows(cows int) *InvalidCowsError {
	if cows > 0 {
		return nil
	}
	if cows < 0 {
		return &InvalidCowsError{cows, "there are no negative cows"}
	}
	return &InvalidCowsError{cows, "no cows don't need food"}
}
