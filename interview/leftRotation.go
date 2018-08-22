package interview

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/teimurjan/go-hackerrank/utils"
)

func runLeftRotation() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	settings := strings.Split(utils.ReadLine(reader), " ")

	arrayLen, err := strconv.ParseInt(settings[0], 10, 64)
	utils.CheckError(err)

	rotationsCount, err := strconv.ParseInt(settings[1], 10, 64)
	utils.CheckError(err)

	rotationsCount %= arrayLen

	arr := strings.Split(utils.ReadLine(reader), " ")
	result := append(arr[rotationsCount:], arr[:rotationsCount]...)

	utils.PrintResult(strings.Join(result, " "), true)
}
