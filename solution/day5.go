package solution

import (
	"bufio"
	"fmt"
	"io"
)

// BF FF BB F RRR
type BoardingPass struct {
	Row    []int
	Column []int
}

type Day5 struct {
	BoardingPasses []BoardingPass
}

func (d *Day5) Input(data io.Reader) error {

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		validated := d.ValidateBSP(line)

		if !validated {
			continue
		}

		boardingPass := d.CreateBoardingPass(line)

		d.BoardingPasses = append(d.BoardingPasses, boardingPass)

	}

	return nil
}

func (d *Day5) ValidateBSP(value string) bool {

	if len(value) != 10 {
		return false
	}

	runes := []rune(value)

	row := runes[:7]
	column := runes[7:]

	if len(row) != 7 || len(column) != 3 {
		return false
	}
	return true
}

func (d *Day5) CreateBoardingPass(bspValue string) BoardingPass {
	runes := []rune(bspValue)

	row := runes[:7]
	column := runes[7:]

	bp := BoardingPass{}

	for _, val := range row {
		var insertVal int

		if string(val) == "F" {
			insertVal = 0
		}

		if string(val) == "B" {
			insertVal = 1
		}

		bp.Row = append(bp.Row, insertVal)
	}

	for _, val := range column {
		var insertVal int

		if string(val) == "L" {
			insertVal = 0
		}

		if string(val) == "R" {
			insertVal = 1
		}
		bp.Column = append(bp.Column, insertVal)
	}

	return bp
}

func (d *Day5) Eliminate(data []int, lower int, upper int) int {

	for _, val := range data {
		diff := (upper - lower) / 2

		switch val {
		case 0:
			upper = upper - diff
		case 1:
			lower = lower + diff
		}
	}

	// (numbered 0 through 127)
	return upper - 1
}

func (d *Day5) Output() {

	var highestSeatNumber int

	for _, pass := range d.BoardingPasses {

		row := d.Eliminate(pass.Row, 0, 128)
		column := d.Eliminate(pass.Column, 0, 8)

		// multiply the row by 8, then add the column.
		// In this example, the seat has ID 44 * 8 + 5 = 357.
		currentSeat := row*8 + column

		if currentSeat > highestSeatNumber {
			highestSeatNumber = currentSeat
		}

	}

	fmt.Printf("=> Answer: The highest seat number is: %d\n", highestSeatNumber)
}
