package day4_test

import (
	"aoc_23/day4"
	"aoc_23/utils"
	"fmt"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	tt := []struct {
		line string
		want int
	}{{line: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", want: 8},
		{line: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", want: 2},
		{line: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", want: 2}}

	for _, tc := range tt {
		t.Run(tc.line, func(t *testing.T) {
			got := day4.CalculatePoints(day4.Card(tc.line))

			if tc.want != got {
				t.Errorf("Want %d, but got %d", tc.want, got)
			}
		})
	}
}

func TestTotalScrachCard(t *testing.T) {
	lines := utils.ReadLines("sample.txt")
	cardCount := make([]int, len(lines))
	for i := range cardCount {
		cardCount[i] = 1
	}

	for i, card := lines{
		
	}
}
