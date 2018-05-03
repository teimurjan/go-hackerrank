package stringss

import (
	"bufio"
	"fmt"
	"os"

	"../utils"
)

func removeSameNeighbours(str string) string {
	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && str[i+1] == str[i] {
			endSlice := ""
			if i+2 < len(str) {
				endSlice = str[i+2:]
			}
			return removeSameNeighbours(str[:i] + endSlice)
		}
	}
	return str
}

func runSuperReducedString() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	str := utils.ReadLine(reader)

	reducedStr := removeSameNeighbours(str)
	if len(reducedStr) == 0 {
		reducedStr = "Empty String"
	}

	utils.PrintResult(fmt.Sprintln(reducedStr), true)
}
