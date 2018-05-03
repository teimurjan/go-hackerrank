package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func PrintResult(result string, debug bool) {
	if !debug {
		outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
		CheckError(err)
		defer outputFile.Close()
		writer := bufio.NewWriterSize(outputFile, 1024*1024)
		fmt.Fprint(writer, result)
		writer.Flush()

	} else {
		fmt.Print(result)
	}
}

func ReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
