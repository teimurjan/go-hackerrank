package warmup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../utils"
)

func runMiniMaxSum() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(utils.ReadLine(reader), " ")

	var arr []int64
	var sum int64
	for _, arrTempItem := range arrTemp {
		arrItem, err := strconv.ParseInt(arrTempItem, 10, 64)
		utils.CheckError(err)

		arr = append(arr, arrItem)
		sum += arrItem
	}

	min := sum
	var max int64
	for _, arrItem := range arr {
		partSum := sum - arrItem
		if partSum > max {
			max = partSum
		}
		if partSum < min {
			min = partSum
		}
	}

	utils.PrintResult(fmt.Sprintf("%d %d\n", min, max), true)
}
