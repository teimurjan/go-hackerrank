package stringss

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"github.com/teimurjan/go-hackerrank/utils"
)

func runCamelCase() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	str := utils.ReadLine(reader)

	wordsCount := 1

	for _, char := range str {
		if unicode.IsUpper(char) {
			wordsCount++
		}
	}

	utils.PrintResult(fmt.Sprintf("%d\n", wordsCount), true)
}
