package main

import (
	"fmt"
	"os"

	"./warmup"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		switch args[0] {
		case "aVeryBigSum":
			warmup.RunAVeryBigSum()
		case "birthdayCakeCandles":
			warmup.RunBirthdayCakeCandles()
		case "diagonalDifference":
			warmup.RunDiagonalDifference()
		case "miniMaxSum":
			warmup.RunMiniMaxSum()
		case "plusMinus":
			warmup.RunPlusMinus()
		case "simpleArraySum":
			warmup.RunSimpleArraySum()
		case "staircase":
			warmup.RunStaircase()
		default:
			fmt.Println("Invalid program name.")
		}
	} else {
		fmt.Println("You need to provide a program to execute(filename without \".go\").")
	}
}
