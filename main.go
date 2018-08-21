package main

import (
	"fmt"
	"os"

	"github.com/teimurjan/go-hackerrank/interview"
	"github.com/teimurjan/go-hackerrank/stringss"
	"github.com/teimurjan/go-hackerrank/warmup"
)

func main() {
	args := os.Args[1:]

	if len(args) == 2 {
		switch args[0] {
		case "warmup":
			warmup.RunSolution(args[1])
		case "strings":
			stringss.RunSolution(args[1])
		case "interview":
			interview.RunSolution(args[1])
		default:
			fmt.Println("Invalid section name.")
		}
	} else {
		fmt.Println("You need to provide exactly 2 args: " +
			"1st is section name(folder name), " +
			"2nd is a program to execute(filename without \".go\") inside this section.")
	}
}
