package warmup

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"../utils"
)

func runStaircase() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	sizeTemp, err := strconv.ParseInt(utils.ReadLine(reader), 10, 64)
	utils.CheckError(err)

	size := int(sizeTemp)

	result := ""
	for i := 0; i < size; i++ {
		sharpsCount := i + 1
		result += strings.Repeat(" ", size-sharpsCount) + strings.Repeat("#", sharpsCount) + "\n"
	}

	utils.PrintResult(result, true)
}
