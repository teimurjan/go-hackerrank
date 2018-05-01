package warmup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunBirthdayCakeCandles() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	readLine(reader)

	candleHeights := strings.Split(readLine(reader), " ")

	var countOf = make(map[int]int)
	var maxHeight int
	for _, candleHeightStr := range candleHeights {
		candleHeightTemp, err := strconv.ParseInt(candleHeightStr, 10, 64)
		checkError(err)
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

	printResult(fmt.Sprintf("%d\n", candlesWillBeBlown), true)
}
