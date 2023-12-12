package day4

import (
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Card string

func CalculatePoints(card Card) int {
	matches := CalculateWinningNumbers(card)
	return int(math.Pow(2, float64(matches-1)))
}

// Calculates no. of winning numbers of current card
func CalculateWinningNumbers(card Card) int {
	_, wnText, hnText := parseCard(card)
	winningNums, havingNums := parseNumbers(wnText), parseNumbers(hnText)
	matches := getMatchingCards(winningNums, havingNums)
	return matches
}

// parse lines into head, winning numbers and having numbers
func parseCard(c Card) (head, wnText, hnText string) {
	sep := strings.Split(string(c), ":")
	head = sep[0]
	sep = strings.Split(sep[1], "|")
	wnText, hnText = sep[0], sep[1]
	return
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

// func
