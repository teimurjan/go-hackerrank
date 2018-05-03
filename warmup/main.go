package warmup

import "fmt"

func RunWarmupSoultion(problemNane string) {
	switch problemNane {
	case "aVeryBigSum":
		runAVeryBigSum()
	case "birthdayCakeCandles":
		runBirthdayCakeCandles()
	case "diagonalDifference":
		runDiagonalDifference()
	case "miniMaxSum":
		runMiniMaxSum()
	case "plusMinus":
		runPlusMinus()
	case "simpleArraySum":
		runSimpleArraySum()
	case "staircase":
		runStaircase()
	case "timeConversion":
		runTimeConversion()
	default:
		fmt.Println("Invalid program name.")
	}
}
