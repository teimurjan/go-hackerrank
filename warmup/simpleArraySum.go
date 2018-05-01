package warmup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func simpleArraySum(arr []int32) int32 {
	var sum int32
	for _, arrItem := range arr {
		sum += arrItem
	}
	return sum
}

func RunSimpleArraySum() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrLength, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arrStr := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(arrLength); i++ {
		arrIntItem, err := strconv.ParseInt(arrStr[i], 10, 64)
		checkError(err)
		arr = append(arr, int32(arrIntItem))
	}

	result := simpleArraySum(arr)

	printResult(fmt.Sprintf("%d\n", result), true)
}
