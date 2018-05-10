package stringss

import "fmt"

func RunStringsSoultion(problemNane string) {
	switch problemNane {
	case "superReducedString":
		runSuperReducedString()
	case "camelCase":
		runCamelCase()
	case "strongPassword":
		runStrongPassword()
	case "caesarCipher":
		runCaesarCipher()
	default:
		fmt.Println("Invalid program name.")
	}
}
