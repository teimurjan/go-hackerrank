package interview

import "fmt"

// RunSolution runs the solution for a problem from interview preparation kit section
func RunSolution(problemNane string) {
	switch problemNane {
	case "2dArrayDS":
		run2dArrayDS()
	case "leftRotation":
		runLeftRotation()
	default:
		fmt.Println("Invalid program name.")
	}
}
