package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ErrInvalidString is Error message.
var ErrInvalidString = errors.New("invalid string")

// Unpack string.
func Unpack(s string) (string, error) {
	var result strings.Builder
	var lastRune rune
	for i, r := range s {
		if unicode.IsDigit(r) && i == 0 {
			return "", ErrInvalidString
		} else if i != len(s)-1 {
			if !unicode.IsDigit(r) && unicode.IsDigit(rune(s[i+1])) {
				if lastRune == r {
					return "", ErrInvalidString
				}
				countRepeat, _ := strconv.Atoi(string(s[i+1]))
				result.WriteString(strings.Repeat(string(r), countRepeat))
				lastRune = r
			} else {
				if !unicode.IsDigit(r) {
					result.WriteString(string(r))
				} else if unicode.IsDigit(rune(s[i+1])) {
					return "", ErrInvalidString
				}
			}
		} else if !unicode.IsDigit(r) {
			result.WriteString(string(r))
			lastRune = r
		}
	}
	return result.String(), nil
}
