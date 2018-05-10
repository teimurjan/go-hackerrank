package stringss

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"../utils"
)

const wordToFind string = "hackerrank"

func isHackerRankInAString(word string) bool {
	passedChars := 0
	for _, char := range word {
		if char == rune(wordToFind[passedChars]) {
			passedChars++
		}
		if len(wordToFind) == passedChars {
			return true
		}
	}
	return false
}

func runHackerRankInAString() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	wordsCount, err := strconv.ParseInt(utils.ReadLine(reader), 10, 64)
	utils.CheckError(err)

	for i := 0; i < int(wordsCount); i++ {
		word := utils.ReadLine(reader)
		result := "NO"
		if isHackerRankInAString(word) {
			result = "YES"
		}
		utils.PrintResult(fmt.Sprintln(result), true)
	}
}
