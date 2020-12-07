package solution

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Day1 struct {
	InputData []int64
	ValueOne  int64
	ValueTwo  int64
}

func (d *Day1) Input(data io.Reader) error {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		valueAsString := scanner.Text()
		converted, err := strconv.ParseInt(valueAsString, 10, 32)

		if err != nil {
			fmt.Println(err)
			return err
		}

		d.InputData = append(d.InputData, converted)
	}

	return nil
}

func (d *Day1) FindCombination() (int64, int64) {

	var value1 int64
	var value2 int64

	for _, initialValue := range d.InputData {
		targetValue := 2020 - initialValue

		for _, value := range d.InputData {
			if value != targetValue || value == initialValue {
				continue
			}

			value1 = initialValue
			value2 = value
		}
	}

	return value1, value2
}

func (d *Day1) Output() string {
	value1, value2 := d.FindCombination()
	multiplied := value1 * value2
	return fmt.Sprintf("COMBINATION IS: %d + %d | ANSWER IS: %d", value1, value2, multiplied)
}
