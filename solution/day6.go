package solution

import (
	"bufio"
	"fmt"
	"io"
)

type Group struct {
	Size    int
	Answers map[string]int
}

func (g *Group) Total() int {
	return len(g.Answers)
}

func (g *Group) Unanimous() int {
	var total int
	for _, amount := range g.Answers {
		if amount != g.Size {
			continue
		}

		total++
	}

	return total
}

type Day6 struct {
	Answers []Group
}

func (d *Day6) Input(data io.Reader) error {
	r := bufio.NewReader(data)

	var group Group

	for {
		line, _, err := r.ReadLine()

		for _, answer := range line {

			if group.Answers == nil {
				group.Answers = make(map[string]int)
			}
			group.Answers[string(answer)]++
		}

		if len(line) == 0 || err == io.EOF {
			d.Answers = append(d.Answers, group)
			group = Group{}
		} else {
			group.Size++
		}

		if err == io.EOF {
			break
		}
	}

	return nil
}

func (d *Day6) Output() {

	var totalSum int
	var unanimousSum int

	for _, group := range d.Answers {
		totalSum += group.Total()
		unanimousSum += group.Unanimous()
	}

	fmt.Printf("=> [PART1] Answer: Sum of answers is %d\n", totalSum)
	fmt.Printf("=> [PART2] Answer: Sum of unanimous answers is %d\n", unanimousSum)
}
