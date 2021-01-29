package converter

import (
	"fmt"
	"math"
	"regexp"
)

// Constants
const (
	Base            = 62
	DigitOffset     = 48
	LowercaseOffset = 61
	UppercaseOffset = 55
)

func ord2char(ord int) (string, error) {
	switch {
	case ord < 10:
		return string(ord + DigitOffset), nil
	case ord >= 10 && ord <= 35:
		return string(ord + UppercaseOffset), nil
	case ord >= 36 && ord < 62:
		return string(ord + LowercaseOffset), nil
	default:
		return "", fmt.Errorf("%d is not a valid integer in the range of base %d", ord, Base)
	}
}

func char2ord(char string) (int, error) {
	if matched, _ := regexp.MatchString("[0-9]", char); matched {
		return int([]rune(char)[0] - DigitOffset), nil
	} else if matched, _ := regexp.MatchString("[A-Z]", char); matched {
		return int([]rune(char)[0] - UppercaseOffset), nil
	} else if matched, _ := regexp.MatchString("[a-z]", char); matched {
		return int([]rune(char)[0] - LowercaseOffset), nil
	} else {
		return -1, fmt.Errorf("%s is not a valid character", char)
	}
}

func reverse(value string) string {
	runes := []rune(value)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes)
}

// Decode ...
func Decode(str string) (int, error) {
	pk := 0
	for i, letter := range reverse(str) {
		d, err := char2ord(string(letter))
		if err != nil {
			return -1, err
		}
		pk = pk + d*int(math.Pow(Base, float64(i)))
	}
	return pk, nil
}

// Encode ...
func Encode(digits int) (string, error) {
	if digits == 0 {
		return "0", nil
	}

	str := ""

	for digits >= 0 {
		remainder := digits % Base
		s, err := ord2char(remainder)
		if err != nil {
			return "", err
		}
		str = s + str
		if digits == 0 {
			break
		}
		digits = digits / Base
	}
	return str, nil
}
