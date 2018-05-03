package warmup

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"

	"../utils"
)

func runAVeryBigSum() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	utils.ReadLine(reader)

	arrStr := strings.Split(utils.ReadLine(reader), " ")

	var sum big.Int
	for _, arrItem := range arrStr {
		arrIntItem := new(big.Int)
		arrIntItem, ok := arrIntItem.SetString(arrItem, 10)
		if !ok {
			panic(arrItem + " is not valid value for big.Int")
		}
		sum.Add(&sum, arrIntItem)
	}

	utils.PrintResult(fmt.Sprintf("%v\n", sum.String()), true)
}
