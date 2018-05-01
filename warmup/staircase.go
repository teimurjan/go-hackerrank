package warmup

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func RunStaircase() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	sizeTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	size := int(sizeTemp)

	result := ""
	for i := 0; i < size; i++ {
		sharpsCount := i + 1
		result += strings.Repeat(" ", size-sharpsCount) + strings.Repeat("#", sharpsCount) + "\n"
	}

	printResult(result, true)
}
