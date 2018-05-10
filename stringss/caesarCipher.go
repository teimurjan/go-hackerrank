package stringss

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"

	"../utils"
)

func getNewLetter(letter rune, rotation int) rune {
	const startLower, startUpper int = int('a'), int('A')
	if unicode.IsUpper(letter) {
		return rune(startUpper + int(math.Mod(float64(int(letter)-startUpper+rotation), 26)))
	} else {
		return rune(startLower + int(math.Mod(float64(int(letter)-startLower+rotation), 26)))
	}
}

func runCaesarCipher() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	_, err := strconv.ParseInt(utils.ReadLine(reader), 10, 64)
	utils.CheckError(err)

	word := utils.ReadLine(reader)

	rotationTemp, err := strconv.ParseInt(utils.ReadLine(reader), 10, 64)
	utils.CheckError(err)

	var encryptedStr string
	rotation := int(math.Mod(float64(rotationTemp), 26))
	for _, char := range word {
		if unicode.IsLetter(char) {
			encryptedStr += string(getNewLetter(char, rotation))
		} else {
			encryptedStr += string(char)
		}
	}

	utils.PrintResult(fmt.Sprintln(encryptedStr), true)
}
