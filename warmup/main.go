package warmup

import "fmt"

// RunSolution runs the solution for a problem from warmup section
func RunSolution(problemNane string) {
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
	case "compareTheTriplets":
		runCompareTheTriplets()
	default:
		fmt.Println("Invalid program name.")
	}
}
