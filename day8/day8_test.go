package day8_test

import (
	"aoc_23/day8"
	"aoc_23/utils"
	"testing"
)

func TestCalculateSteps(t *testing.T) {
	lines := utils.ReadLines("sample.txt")
	got := day8.CalculateSteps(lines)
	want := 2

	if got != want {
		t.Errorf("got %d, but want %d", got, want)
	}
}
