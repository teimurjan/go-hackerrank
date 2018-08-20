package stringss

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"github.com/teimurjan/go-hackerrank/utils"
)

func isSpecialChar(char rune) bool {
	specialCharacters := "!@#$%^&*()-+"
	for _, specialChar := range specialCharacters {
		if specialChar == char {
			return true
		}
	}
	return false
}

func countLackingChars(word string) int {
	var isNumberPresent, isLowerPresent, isUpperPresent, isSpecialPresent bool

	lackingCharsCount := 4

	for _, char := range word {
		if unicode.IsNumber(char) && !isNumberPresent {
			isNumberPresent = true
			lackingCharsCount--
		} else if unicode.IsLower(char) && !isLowerPresent {
			isLowerPresent = true
			lackingCharsCount--
		} else if unicode.IsUpper(char) && !isUpperPresent {
			isUpperPresent = true
			lackingCharsCount--
		} else if isSpecialChar(char) && !isSpecialPresent {
			isSpecialPresent = true
			lackingCharsCount--
		}
	}

	return lackingCharsCount
}

const minPasswordLength = 6

func runStrongPassword() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	utils.ReadLine(reader)
	password := utils.ReadLine(reader)

	minCharsToMakePasswordStrong := countLackingChars(password)

	if minPasswordLength-len(password) > minCharsToMakePasswordStrong {
		minCharsToMakePasswordStrong = minPasswordLength - len(password)
	}
	utils.PrintResult(fmt.Sprintf("%d\n", minCharsToMakePasswordStrong), true)
}
