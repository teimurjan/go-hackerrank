package warmup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"../utils"
)

func runTimeConversion() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	time := utils.ReadLine(reader)

	var result string
	if hoursTemp, isPM := time[:2], time[len(time)-2:] == "PM"; hoursTemp == "12" && !isPM {
		result = "00" + time[2:len(time)-2] + "\n"
	} else if isPM && hoursTemp != "12" {
		var hoursStr string
		if hoursTemp[0] == '0' {
			hoursStr = string(hoursTemp[1])
		} else {
			hoursStr = hoursTemp
		}

		hours, err := strconv.ParseInt(hoursStr, 10, 64)
		utils.CheckError(err)

		newHours := hours + 12
		result = fmt.Sprintf("%d%v\n", newHours, time[2:len(time)-2])
	} else {
		result = time[:len(time)-2] + "\n"
	}

	utils.PrintResult(result, true)
}
