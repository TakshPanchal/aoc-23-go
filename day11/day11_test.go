package day11_test

import (
	"aoc_23/day11"
	"aoc_23/utils"
	"fmt"
	"math"
	"slices"
	"testing"
)

func TestFindClosestDistance(t *testing.T) {
	tt := []struct {
		g1, g2 day11.Galaxy
		want   int
	}{
		{g1: day11.Galaxy{2, 7}, g2: day11.Galaxy{6, 12}, want: 9},
	}

	for i, tc := range tt {
		t.Run(string(rune(i)), func(t *testing.T) {
			got := day11.FindClosestDistance(tc.g1, tc.g2)

			if got != tc.want {
				t.Errorf("got %d, but want %d", got, tc.want)
			}
		})
	}
}

func TestPar2(t *testing.T) {
	const EXP_VALUE = 1000000
	img := utils.ReadLines("sample.txt")
	idx := day11.GetEmptyRowColIdx(img)
	rIdx, cIdx := idx[0], idx[1]
	slices.Sort(rIdx)
	slices.Sort(cIdx)

	fmt.Println(idx[0])
	fmt.Println(idx[1])

	gs := day11.FindGalaxies(img)
	sum := 0
	for i := range gs {
		for j := i + 1; j < len(gs); j++ {
			sum += day11.FindClosestDistance(gs[i], gs[j])
			xMin, xMax := int(math.Min(float64(gs[i].X), float64(gs[j].X))),
				int(math.Max(float64(gs[i].X), float64(gs[j].X)))
			yMin, yMax := int(math.Min(float64(gs[i].Y), float64(gs[j].Y))),
				int(math.Max(float64(gs[i].Y), float64(gs[j].Y)))

			for _, r := range rIdx {
				if xMin-1 <= r && r <= xMax-1 {
					sum += EXP_VALUE - 1
				}
			}

			for _, c := range cIdx {
				if yMin-1 <= c && c <= yMax-1 {
					sum += EXP_VALUE - 1
				}
			}

		}
	}
	fmt.Println(sum)
}
