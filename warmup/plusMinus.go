package warmup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunPlusMinus() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)
	defer writer.Flush()

	arrSizeTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arrTemp := strings.Split(readLine(reader), " ")

	var positivesCount, negativesCount, zerosCount int32

	for _, arrTempItem := range arrTemp {
		arrItem, err := strconv.ParseInt(arrTempItem, 10, 64)
		checkError(err)

		switch {
		case arrItem == 0:
			zerosCount++
		case arrItem > 0:
			positivesCount++
		default:
			negativesCount++
		}
	}

	printResult(fmt.Sprintf("%f\n", float64(positivesCount)/float64(arrSizeTemp)), true)
	printResult(fmt.Sprintf("%f\n", float64(negativesCount)/float64(arrSizeTemp)), true)
	printResult(fmt.Sprintf("%f\n", float64(zerosCount)/float64(arrSizeTemp)), true)
}
