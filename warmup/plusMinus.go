package warmup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../utils"
)

func runPlusMinus() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	utils.CheckError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)
	defer writer.Flush()

	arrSizeTemp, err := strconv.ParseInt(utils.ReadLine(reader), 10, 64)
	utils.CheckError(err)

	arrTemp := strings.Split(utils.ReadLine(reader), " ")

	var positivesCount, negativesCount, zerosCount int32

	for _, arrTempItem := range arrTemp {
		arrItem, err := strconv.ParseInt(arrTempItem, 10, 64)
		utils.CheckError(err)

		switch {
		case arrItem == 0:
			zerosCount++
		case arrItem > 0:
			positivesCount++
		default:
			negativesCount++
		}
	}

	utils.PrintResult(fmt.Sprintf("%f\n", float64(positivesCount)/float64(arrSizeTemp)), true)
	utils.PrintResult(fmt.Sprintf("%f\n", float64(negativesCount)/float64(arrSizeTemp)), true)
	utils.PrintResult(fmt.Sprintf("%f\n", float64(zerosCount)/float64(arrSizeTemp)), true)
}
