package warmup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/teimurjan/go-hackerrank/utils"
)

func runBirthdayCakeCandles() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	utils.ReadLine(reader)

	candleHeights := strings.Split(utils.ReadLine(reader), " ")

	var countOf = make(map[int]int)
	var maxHeight int
	for _, candleHeightStr := range candleHeights {
		candleHeightTemp, err := strconv.ParseInt(candleHeightStr, 10, 64)
		utils.CheckError(err)
		candleHeight := int(candleHeightTemp)

		if currentCandleHeight := int(candleHeight); currentCandleHeight > maxHeight {
			maxHeight = currentCandleHeight
		}

		if _, ok := countOf[candleHeight]; ok {
			countOf[candleHeight]++
		} else {
			countOf[candleHeight] = 1
		}
	}

	candlesWillBeBlown := countOf[maxHeight]

	utils.PrintResult(fmt.Sprintf("%d\n", candlesWillBeBlown), true)
}
