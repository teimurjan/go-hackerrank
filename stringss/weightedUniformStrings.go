package stringss

import (
	"bufio"
	"os"
	"strconv"
	"unicode"

	"../utils"
)

func getWeights(str *string) map[int]bool {
	weights := make(map[int]bool)

	const startLower, startUpper int = int('a'), int('A')

	var prevCh rune
	var series int
	for _, ch := range *str {
		alphabetStart := startLower
		if unicode.IsUpper(ch) {
			alphabetStart = startUpper
		}

		if prevCh != ch {
			series = 1
		} else {
			series++
		}

		weights[(int(ch)-alphabetStart+1)*series] = true
		prevCh = ch
	}

	return weights
}

func runWeightedUniformStrings() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	str := utils.ReadLine(reader)
	weights := getWeights(&str)

	inputCount, err := strconv.ParseInt(utils.ReadLine(reader), 10, 64)
	utils.CheckError(err)

	results := ""
	for i := 0; i < int(inputCount); i++ {
		queryTemp := utils.ReadLine(reader)
		query, err := strconv.ParseInt(queryTemp, 10, 64)
		utils.CheckError(err)

		result := "Yes"
		_, ok := weights[int(query)]
		if !ok {
			result = "No"
		}
		results += result + "\n"
	}

	utils.PrintResult(results, true)
}
