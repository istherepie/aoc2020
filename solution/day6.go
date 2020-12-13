package solution

import (
	"bufio"
	"fmt"
	"io"
)

type Day6 struct {
	Answers []map[string]struct{}
}

func (d *Day6) Input(data io.Reader) error {
	r := bufio.NewReader(data)

	var group map[string]struct{}

	for {
		line, _, err := r.ReadLine()

		if group == nil {
			group = make(map[string]struct{})
		}

		for _, answer := range line {
			group[string(answer)] = Empty{}
		}

		if len(line) == 0 || err == io.EOF {
			d.Answers = append(d.Answers, group)
			group = nil
		}

		if err == io.EOF {
			break
		}
	}

	return nil
}

func (d *Day6) Output() {

	var sumOfAnswers int

	for _, answer := range d.Answers {
		sumOfAnswers += len(answer)
	}

	fmt.Printf("=> Answer: Sum of answers is %d\n", sumOfAnswers)
}
