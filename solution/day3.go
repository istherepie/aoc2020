package solution

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Slope struct {
	Right int
	Down  int
}

type Day3 struct {
	Patterns []string
}

func (d *Day3) Input(data io.Reader) error {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		d.Patterns = append(d.Patterns, line)
	}

	return nil
}

func (d *Day3) FindPath(right int, down int) int {

	var count int
	var position int

	// iteration
	i := 0

	for {

		// Move 3 to the right
		position = position + right

		// Move 1 down
		i = i + down

		// Finish when last pattern reached
		if i >= len(d.Patterns) {
			break
		}

		pattern := d.Patterns[i]

		if position >= len(pattern) {

			patternRepeat := position/len(pattern) + 1
			pattern = strings.Repeat(pattern, patternRepeat)

		}

		if string(pattern[position]) == "#" {
			count++
		}

	}

	return count
}

func (d *Day3) Output() {

	// Part 1
	treesEncountered := d.FindPath(3, 1)

	// Part 2
	var results []int

	slopes := []Slope{
		Slope{1, 1},
		Slope{3, 1},
		Slope{5, 1},
		Slope{7, 1},
		Slope{1, 2},
	}

	for _, slope := range slopes {
		treesEncountered := d.FindPath(slope.Right, slope.Down)
		results = append(results, treesEncountered)
	}

	var treesMultiplied int
	for _, result := range results {
		if treesMultiplied <= 0 {
			treesMultiplied = result
			continue
		}

		treesMultiplied = treesMultiplied * result
	}

	fmt.Printf("=> [PART1] Answer: %d trees encountered\n", treesEncountered)
	fmt.Printf("=> [PART2] Answer: %d (Trees encountered: %d)\n", treesMultiplied, results)
}
