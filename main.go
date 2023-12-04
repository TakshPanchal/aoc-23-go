package main

import (
	"aoc_23/day2"
	"fmt"
	"os"
	"strings"
)

type Line string

func main() {
	fData, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// fData to list of Lines
	ls := strings.Split(string(fData), "\n")

	/// Day 1

	// // sum := day1.SumCalibration(ls)
	// sum := 0
	// for _, l := range ls {
	// 	sum += day1.Calibration(l)
	// }

	// fmt.Printf("sum of all of the calibration values is %d \n", sum)

	/* Day 2 */
	sum := 0
	bag := day2.Bag{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	for _, l := range ls {
		g, err := day2.ExtractGame(l)
		if err != nil {
			panic(err)
		}

		if day2.IsPossible(bag, g) {
			sum += g.Id
		}
	}

	fmt.Printf("sum of the IDs of possible games %d \n", sum)
}
