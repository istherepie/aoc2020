package solution

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

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

func (d *Day3) FindPath() int {

	var count int
	var position int

	// Skip first
	for _, pattern := range d.Patterns[1:] {

		// Move 3 to the right
		position = position + 3

		if position > len(pattern) {

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
	treesEncountered := d.FindPath()
	fmt.Printf("=> Answer: %d trees encountered\n", treesEncountered)
}
