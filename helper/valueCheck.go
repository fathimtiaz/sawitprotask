package helper

import "unicode"

func StringHasNumber(str string) bool {
	for _, char := range str {
		if unicode.IsDigit(char) {
			return true
		}
	}

	return false
}

func StringHasCapital(str string) bool {
	for _, char := range str {
		if unicode.IsUpper(char) {
			return true
		}
	}

	return false
}

func StringHasSpecialChar(str string) bool {
	for _, char := range str {
		if unicode.IsSymbol(char) {
			return true
		}
	}

	return false
}
