package phonenumber

import (
	"errors"
	"fmt"
)

func filterDigits(s string) []rune {
	digits := make([]rune, 0)
	for _, c := range s {
		if '0' <= c && c <= '9' {
			digits = append(digits, c)
		}
	}
	return digits
}

func Number(phoneNumber string) (string, error) {
	digits := filterDigits(phoneNumber)

	if len(digits) == 11 { // with country code
		if digits[0] == '1' {
			digits = digits[1:] // drop valid country code
		} else {
			return "", errors.New("Invalid number: " + phoneNumber)
		}
	}

	if len(digits) != 10 {
		return "", errors.New("Invalid number: " + phoneNumber)
	}

	if digits[0] < '2' || digits[3] < '2' {
		return "", errors.New("Invalid number: " + phoneNumber)
	}

	return string(digits), nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return phoneNumber, err
	}
	return phoneNumber[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return phoneNumber, err
	}
	areaCode, err := AreaCode(phoneNumber)
	if err != nil {
		return phoneNumber, err
	}
	exchangeCode := phoneNumber[3:6]
	subscriberNumber := phoneNumber[6:]
	return fmt.Sprintf("(%s) %s-%s", areaCode, exchangeCode, subscriberNumber), nil
}
