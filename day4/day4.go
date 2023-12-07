package day4

import (
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func CalculatePoints(text string) int {
	sep := strings.Split(text, ":")
	numsText := sep[1]
	sep = strings.Split(numsText, "|")
	wnText, hnText := sep[0], sep[1]

	winningNums, havingNums := parseNumbers(wnText), parseNumbers(hnText)
	matches := getMatchingCards(winningNums, havingNums)
	return int(math.Pow(2, float64(matches-1)))
}

// parse list of numbers in string seprated by spaces
func parseNumbers(list string) []int {
	list = strings.Trim(list, " ")
	nums := []int{}
	nTxt := strings.Split(list, " ")
	for i := range nTxt {
		if len(nTxt[i]) != 0 {
			// TODO: Deal with the error
			n, _ := strconv.Atoi(nTxt[i])
			nums = append(nums, n)
		}
	}
	return nums
}

func getMatchingCards(winningNums []int, havingNums []int) int {
	matches := 0
	slices.Sort(havingNums)
	for _, w := range winningNums {
		idx := sort.SearchInts(havingNums, w)
		if idx < len(havingNums) && havingNums[idx] == w {
			matches++
		}
	}
	return matches
}
