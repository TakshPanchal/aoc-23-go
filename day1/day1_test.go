package day1_test

import (
	"aoc_23/day1"
	"testing"
)

func TestCalibration(t *testing.T) {
	tests := []struct {
		line string
		want int
	}{
		{line: "eightwothree", want: 83},
		{line: "7pqrstsixteen", want: 76},
		{line: "24", want: 24},
		{line: "eightwo3", want: 83},
		{line: "eight", want: 88},
		{line: "eighttwo", want: 82},
		{line: "zoneight234", want: 14},
		{line: "three7two", want: 32},
		{line: "1one", want: 11},
		{line: "zoneight234", want: 14},
		{line: "three7twothree", want: 33},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			got := day1.Calibration(tt.line)

			if got != tt.want {
				t.Errorf("got: %d, but want %d", got, tt.want)
			}
		})
	}
}
