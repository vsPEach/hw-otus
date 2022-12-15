package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	res := new(strings.Builder)
	input := []rune(str)

	if len(input) == 0 {
		return "", nil
	}

	if isNum(input[0]) {
		return "", ErrInvalidString
	}

	for i := 0; i < len(input)-1; i++ {
		if isNum(input[i]) && isNum(input[i+1]) {
			return "", ErrInvalidString
		}
		if !isNum(input[i]) && isNum(input[i+1]) {
			count, err := strconv.Atoi(string(input[i+1]))
			if err != nil {
				return "", err
			}
			substr := strings.Repeat(string(input[i]), count)
			res.WriteString(substr)
			continue
		} else if !isNum(input[i]) {
			res.WriteString(string(input[i]))
		}
	}
	if !isNum(input[len(input)-1]) {
		res.WriteRune(input[len(input)-1])
	}
	return res.String(), nil
}

func isNum(num rune) bool {
	return (num <= 57) && (num >= 48)
}
