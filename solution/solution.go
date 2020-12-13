package solution

import (
	"io"
)

type Solution interface {
	Input(io.Reader) error
	Output()
}

var All map[string]Solution = map[string]Solution{
	"day1": &Day1{},
	"day2": &Day2{},
	"day3": &Day3{},
	"day4": &Day4{},
	"day5": &Day5{},
	"day6": &Day6{},
}
