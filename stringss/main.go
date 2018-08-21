package stringss

import "fmt"

// RunSolution runs the solution for a problem from strings section
func RunSolution(problemNane string) {
	switch problemNane {
	case "superReducedString":
		runSuperReducedString()
	case "camelCase":
		runCamelCase()
	case "strongPassword":
		runStrongPassword()
	case "caesarCipher":
		runCaesarCipher()
	case "marsExploration":
		runMarsExploration()
	case "hackerRankInAString":
		runHackerRankInAString()
	case "pangrams":
		runPangrams()
	case "weightedUniformStrings":
		runWeightedUniformStrings()
	default:
		fmt.Println("Invalid program name.")
	}
}
