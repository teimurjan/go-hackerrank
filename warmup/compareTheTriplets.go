package warmup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../utils"
)

func parseTripletStr(tripletStr string) [3]int {
	tripletArrStr := strings.Split(tripletStr, " ")
	var tripletArr [3]int
	for i, itemStr := range tripletArrStr {
		item, err := strconv.ParseInt(itemStr, 10, 64)
		utils.CheckError(err)
		tripletArr[i] = int(item)
	}
	return tripletArr
}

func compareTriplets(triplet1 [3]int, triplet2 [3]int) [2]int {
	var points [2]int
	for i := 0; i < 3; i++ {
		triplet1Item := triplet1[i]
		triplet2Item := triplet2[i]
		if triplet1Item > triplet2Item {
			points[0]++
		} else if triplet2Item > triplet1Item {
			points[1]++
		}
	}
	return points
}

func runCompareTheTriplets() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	aliceTripletStr := utils.ReadLine(reader)
	bobTripletStr := utils.ReadLine(reader)

	aliceTriplet := parseTripletStr(aliceTripletStr)
	bobTriplet := parseTripletStr(bobTripletStr)

	points := compareTriplets(aliceTriplet, bobTriplet)

	utils.PrintResult(fmt.Sprintf("%d %d\n", points[0], points[1]), true)
}
