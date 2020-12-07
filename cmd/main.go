package main

import (
	"fmt"
	"os"

	"github.com/istherepie/aoc2020/solution"
)

func main() {
	fileContents, err := os.Open("inputdata/day1.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := solution.Day1{}

	inputErr := s.Input(fileContents)
	fileContents.Close()

	if inputErr != nil {
		fmt.Println(inputErr)
		os.Exit(1)
	}

	s.Output()

	fileContents, err = os.Open("inputdata/day2.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s2 := solution.Day2{}

	inputErr = s2.Input(fileContents)
	fileContents.Close()

	if inputErr != nil {
		fmt.Println(inputErr)
		os.Exit(1)
	}

	s2.Output()
}
