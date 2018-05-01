package warmup

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func RunDiagonalDifference() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	matrixSize64, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	matrixSize := int32(matrixSize64)

	var primaryDiagonalSum, secondaryDiagonalSum int32

	for i := 0; i < int(matrixSize); i++ {
		rowStr := strings.Split(readLine(reader), " ")

		primaryI, err := strconv.ParseInt(rowStr[i], 10, 64)
		checkError(err)
		primaryDiagonalSum += int32(primaryI)

		secondaryI, err := strconv.ParseInt(rowStr[(matrixSize-1)-int32(i)], 10, 64)
		checkError(err)
		secondaryDiagonalSum += int32(secondaryI)
	}

	result := math.Abs(float64(primaryDiagonalSum - secondaryDiagonalSum))
	printResult(fmt.Sprintf("%d\n", int32(result)), true)
}
