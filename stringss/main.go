package stringss

import "fmt"

func RunStringsSoultion(problemNane string) {
	switch problemNane {
	case "superReducedString":
		runSuperReducedString()
	default:
		fmt.Println("Invalid program name.")
	}
}
