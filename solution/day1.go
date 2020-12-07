package solution

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Day1 struct {
	InputData []int64
}

func (d *Day1) Input(data io.Reader) error {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		valueAsString := scanner.Text()
		value64, err := strconv.ParseInt(valueAsString, 10, 32)

		if err != nil {
			fmt.Println(err)
			return err
		}

		d.InputData = append(d.InputData, value64)
	}

	return nil
}

func (d *Day1) FindTargetValue(initial int64, target int64) (int64, bool) {

	var found bool
	var match int64

	for _, value := range d.InputData {

		if value == initial {
			continue
		}

		if value == target {
			found = true
			match = value
		}
	}

	return match, found
}

func (d *Day1) FindCombinationPart1() (int64, int64) {

	var value1 int64
	var value2 int64

	for _, initialValue := range d.InputData {
		targetValue := 2020 - initialValue

		match, found := d.FindTargetValue(initialValue, targetValue)

		if !found {
			continue
		}

		value1 = initialValue
		value2 = match

	}

	return value1, value2
}

func (d *Day1) FindCombinationPart2() (int64, int64, int64) {

	var value1 int64
	var value2 int64
	var value3 int64

	for _, initialValue := range d.InputData {
		targetValue := 2020 - initialValue

		for _, nextValue := range d.InputData {

			newTarget := targetValue - nextValue

			if newTarget < 0 {
				continue
			}

			match, found := d.FindTargetValue(nextValue, newTarget)

			if !found {
				continue
			}

			value1 = initialValue
			value2 = nextValue
			value3 = match
		}
	}

	return value1, value2, value3
}

func (d *Day1) Output() {
	value1, value2 := d.FindCombinationPart1()
	multiplied := value1 * value2
	fmt.Printf("=> [PART1] Answer: %d (combination: %d + %d)\n", multiplied, value1, value2)

	value1, value2, value3 := d.FindCombinationPart2()
	multiplied = value1 * value2 * value3
	fmt.Printf("=> [PART2] Answer: %d (combination: %d + %d + %d)\n", multiplied, value1, value2, value3)
}
