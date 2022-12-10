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
	if len(input) != 0 && unicode.IsDigit(input[0]) {
		return "", ErrInvalidString
	}
	if str == "" {
		return "", nil
	}
	for i := 0; i < len(input)-1; i++ {
		if unicode.IsDigit(input[i]) && unicode.IsDigit(input[i+1]) {
			return "", ErrInvalidString
		}
		if !unicode.IsDigit(input[i]) && unicode.IsDigit(input[i+1]) {
			count, err := strconv.Atoi(string(input[i+1]))
			if err != nil {
				return "", err
			}
			substr := strings.Repeat(string(input[i]), count)
			res.WriteString(substr)
			continue
		} else if !unicode.IsDigit(input[i]) {
			res.WriteString(string(input[i]))
		}
	}
	if !unicode.IsDigit(input[len(input)-1]) {
		res.WriteRune(input[len(input)-1])
	}
	return res.String(), nil
}
