package stringss

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/teimurjan/go-hackerrank/utils"
)

const alphabetStartIndex int = 97
const alphabetEndIndex int = 122

func getAlphabet() []rune {
	var alphabet []rune
	for i := alphabetStartIndex; i < alphabetEndIndex+1; i++ {
		alphabet = append(alphabet, rune(i))
	}
	return alphabet
}

func areLettersEqual(a rune, b rune) bool {
	return strings.ToLower(string(a)) == strings.ToLower(string(b))
}

func isPangram(s string) bool {
	alphabet := getAlphabet()

	lacksCount := len(alphabet)

	for _, char := range s {
		for i, alphabetChar := range alphabet {
			if areLettersEqual(char, alphabetChar) {
				alphabet = append(alphabet[:i], alphabet[i+1:]...)
				lacksCount--
				if lacksCount == 0 {
					return true
				}
			}
		}
	}

	return false
}

func runPangrams() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	str := utils.ReadLine(reader)

	result := "not pangram"
	if isPangram(str) {
		result = "pangram"
	}

	utils.PrintResult(fmt.Sprintln(result), true)
}
