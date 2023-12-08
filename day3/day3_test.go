package day3_test

import (
	"aoc_23/day3"
	"aoc_23/utils"
	"reflect"
	"testing"
)

func TestFindPartNumbers(t *testing.T) {
	engineSchematic := utils.ReadLines("sample.txt")
	got := day3.FindPartNumbers(engineSchematic)
	want := []int{467, 35, 633, 617, 592, 755, 664, 598}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestFindGears(t *testing.T) {
	engineSchematic := utils.ReadLines("sample.txt")
	got := day3.FindGears(engineSchematic)
	want := []day3.Gear{
		{N1: 467, N2: 35},
		{755, 598},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}

}
