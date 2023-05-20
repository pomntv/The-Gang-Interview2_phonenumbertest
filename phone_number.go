package phonenumber

import (
	"errors"
	"regexp"
)

// Number cleans and validates a phone number.
func Number(input string) (string, error) {
	// Remove all non-digits from the input
	re := regexp.MustCompile(`\D`)
	stripped := re.ReplaceAllString(input, "")

	// If the number is 11 digits and starts with a 1, strip the 1
	if len(stripped) == 11 && stripped[0] == '1' {
		stripped = stripped[1:]
	}

	// Check if the number has the correct length
	if len(stripped) != 10 {
		return "", errors.New("incorrect number of digits")
	}

	// Check if the area code or exchange code start with 0 or 1
	if stripped[0] == '0' || stripped[0] == '1' || stripped[3] == '0' || stripped[3] == '1' {
		return "", errors.New("area code or exchange code cannot start with zero or one")
	}

	return stripped, nil
}

// AreaCode extracts the area code from a phone number.
func AreaCode(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}

	return number[:3], nil
}

// Format formats a phone number into the standard (XXX) XXX-XXXX format.
func Format(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}

	// Format the number
	formatted := "(" + number[:3] + ") " + number[3:6] + "-" + number[6:]

	return formatted, nil
}
