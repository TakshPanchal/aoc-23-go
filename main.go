package main

import (
	"aoc_23/day4"
	"aoc_23/utils"
	"fmt"
)

type Line string

func main() {
	ls := utils.ReadLines("input.txt")

	/// Day 1

	// // sum := day1.SumCalibration(ls)
	// sum := 0
	// for _, l := range ls {
	// 	sum += day1.Calibration(l)
	// }

	// fmt.Printf("sum of all of the calibration values is %d \n", sum)

	/* Day 2 */

	// part 1
	// sum := 0
	// bag := day2.Bag{
	// 	Red:   12,
	// 	Green: 13,
	// 	Blue:  14,
	// }

	// for _, l := range ls {
	// 	g, err := day2.ExtractGame(l)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if day2.IsPossible(bag, g) {
	// 		sum += g.Id
	// 	}
	// }

	// fmt.Printf("sum of the IDs of possible games %d \n", sum)

	// // part2
	// sum = 0
	// for _, l := range ls {
	// 	g, err := day2.ExtractGame(l)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	b := day2.FewestPossibleBag(g)

	// 	sum += b.Red * b.Green * b.Blue
	// }
	// fmt.Printf("sum of the power of these fewest number of cubes of each color %d \n", sum)

	/*---------------------------------------------------------*/

	/* --- Day 3: Gear Ratios --- */

	// part1
	// partNums := day3.FindPartNumbers(ls)
	// sum := utils.SumAll(partNums)
	// fmt.Printf("sum of all of the part numbers in the engine schematic %d \n", sum)

	/* --- Day 4: Scratchcards --- */

	// Part1
	points := 0
	for _, l := range ls {
		points += day4.CalculatePoints(l)
	}
	fmt.Printf("No. of points are they worth in total is c%d \n", points)
}
