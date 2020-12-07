package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/istherepie/aoc2020/solution"
)

func main() {

	pathToInputData := "inputdata"
	for c, s := range solution.All {

		filename := c + ".txt"
		file := filepath.Join(pathToInputData, filename)

		fileContents, err := os.Open(file)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		inputErr := s.Input(fileContents)
		fileContents.Close()

		if inputErr != nil {
			fmt.Println(inputErr)
			os.Exit(1)
		}

		fmt.Printf("Challenge: %v\n", c)
		s.Output()

		// Padding...
		fmt.Println()
	}
}
