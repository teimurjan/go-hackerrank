package interview

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/teimurjan/go-hackerrank/utils"
)

func parseRowToArr(row string) [6]int64 {
	arr := [6]int64{}
	arrRow := strings.Split(row, " ")
	for i, strValue := range arrRow {
		value, err := strconv.ParseInt(strValue, 10, 64)
		utils.CheckError(err)
		arr[i] = value
	}
	return arr
}

func getMaxHourglassSum(matrix [6][6]int64) int64 {
	var maxSum int64 = -60
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			firstRowSum := matrix[i][j] + matrix[i][j+1] + matrix[i][j+2]
			secondRowSum := matrix[i+1][j+1]
			thirdRowSum := matrix[i+2][j] + matrix[i+2][j+1] + matrix[i+2][j+2]
			sum := firstRowSum + secondRowSum + thirdRowSum
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func run2dArrayDS() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	matrix := [6][6]int64{}
	for i := 0; i < 6; i++ {
		arrRowTemp := utils.ReadLine(reader)
		matrix[i] = parseRowToArr(arrRowTemp)
	}

	utils.PrintResult(fmt.Sprintln(getMaxHourglassSum(matrix)), true)
}
