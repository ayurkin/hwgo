package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedStr string) (string, error) {
	var sb strings.Builder
	packedStrSlice := []rune(packedStr)
	prevSym := rune('0')

	for i, curSym := range packedStrSlice {
		isDigitPrevSym := unicode.IsDigit(prevSym)
		isDigitCurSym := unicode.IsDigit(curSym)

		switch {
		case isDigitPrevSym && isDigitCurSym:
			return "", ErrInvalidString
		case !isDigitPrevSym && isDigitCurSym:
			repeatNumber, _ := strconv.Atoi(string(curSym))
			sb.WriteString(strings.Repeat(string(prevSym), repeatNumber))
		case !isDigitPrevSym && !isDigitCurSym:
			sb.WriteString(string(prevSym))
		}

		if !isDigitCurSym && i == len(packedStrSlice)-1 {
			sb.WriteString(string(curSym))
			break
		}

		prevSym = curSym
	}

	return sb.String(), nil
}
