package warmup

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/teimurjan/go-hackerrank/utils"
)

func runDiagonalDifference() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	matrixSize64, err := strconv.ParseInt(utils.ReadLine(reader), 10, 64)
	utils.CheckError(err)

	matrixSize := int32(matrixSize64)

	var primaryDiagonalSum, secondaryDiagonalSum int32

	for i := 0; i < int(matrixSize); i++ {
		rowStr := strings.Split(utils.ReadLine(reader), " ")

		primaryI, err := strconv.ParseInt(rowStr[i], 10, 64)
		utils.CheckError(err)
		primaryDiagonalSum += int32(primaryI)

		secondaryI, err := strconv.ParseInt(rowStr[(matrixSize-1)-int32(i)], 10, 64)
		utils.CheckError(err)
		secondaryDiagonalSum += int32(secondaryI)
	}

	result := math.Abs(float64(primaryDiagonalSum - secondaryDiagonalSum))
	utils.PrintResult(fmt.Sprintf("%d\n", int32(result)), true)
}
