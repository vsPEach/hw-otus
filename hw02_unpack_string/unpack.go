package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	res := new(strings.Builder)
	input := []rune(str)

	if len(input) == 0 {
		return "", nil
	}

	if unicode.IsDigit(input[0]) {
		return "", ErrInvalidString
	}

	for i := 0; i < len(input)-1; i++ {
		if ((input[i] <= 57) && (input[i] >= 48)) && (input[i+1] <= 57) && (input[i+1] >= 48) {
			return "", ErrInvalidString
		}
		if !((input[i] <= 57) && (input[i] >= 48)) && ((input[i+1] <= 57) && (input[i+1] >= 48)) {
			count, err := strconv.Atoi(string(input[i+1]))
			if err != nil {
				return "", err
			}
			substr := strings.Repeat(string(input[i]), count)
			res.WriteString(substr)
			continue
		} else if !(48 <= input[i] && input[i] <= 57) {
			res.WriteString(string(input[i]))
		}
	}
	if !(48 <= input[len(input)-1] && input[len(input)-1] <= 57) {
		res.WriteRune(input[len(input)-1])
	}
	return res.String(), nil
}
