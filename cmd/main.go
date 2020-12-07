package main

import (
	"fmt"
	"os"

	"github.com/istherepie/aoc2020/solution"
)

func main() {
	fileContents, err := os.Open("inputdata/day1.txt")
	defer fileContents.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := solution.Day1{}

	inputErr := s.Input(fileContents)

	if inputErr != nil {
		fmt.Println(inputErr)
		os.Exit(1)
	}

	output := s.Output()

	fmt.Println(output)
}
