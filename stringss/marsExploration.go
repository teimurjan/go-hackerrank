package stringss

import (
	"bufio"
	"fmt"
	"os"

	"github.com/teimurjan/go-hackerrank/utils"
)

const originalMsg string = "SOS"

func countErrors(sosMsg string) int {
	errorsCount := 0

	for i, char := range sosMsg {
		if char != rune(originalMsg[i]) {
			errorsCount++
		}
	}

	return errorsCount
}

func runMarsExploration() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	str := utils.ReadLine(reader)

	errorsCount := 0
	for i := 3; i <= len(str); i += 3 {
		errorsCount += countErrors(str[i-3 : i])
	}

	utils.PrintResult(fmt.Sprintln(errorsCount), true)
}
