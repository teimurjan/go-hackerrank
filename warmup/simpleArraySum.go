package warmup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../utils"
)

func simpleArraySum(arr []int32) int32 {
	var sum int32
	for _, arrItem := range arr {
		sum += arrItem
	}
	return sum
}

func runSimpleArraySum() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrLength, err := strconv.ParseInt(utils.ReadLine(reader), 10, 64)
	utils.CheckError(err)

	arrStr := strings.Split(utils.ReadLine(reader), " ")

	var arr []int32

	for i := 0; i < int(arrLength); i++ {
		arrIntItem, err := strconv.ParseInt(arrStr[i], 10, 64)
		utils.CheckError(err)
		arr = append(arr, int32(arrIntItem))
	}

	result := simpleArraySum(arr)

	utils.PrintResult(fmt.Sprintf("%d\n", result), true)
}
