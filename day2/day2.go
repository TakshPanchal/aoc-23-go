// --- Day 2: Cube Conundrum ---
package day2

import (
	"errors"
	"strconv"
	"strings"
)

type ColorSet struct {
	Red, Green, Blue int
}

type Bag ColorSet

type Game struct {
	Id      int
	Reveals []ColorSet
}

// Given bag and a game function will return weather the game is possible or not,
func IsPossible(bag Bag, g Game) bool {
	for _, r := range g.Reveals {
		if r.Blue > bag.Blue ||
			r.Red > bag.Red ||
			r.Green > bag.Green {
			return false
		}
	}
	return true
}

var ErrInput = errors.New("something wrong in input format")

func ExtractGame(l string) (Game, error) {
	s := strings.Split(l, ":")

	if len(s) != 2 {
		return Game{}, ErrInput
	}

	hd, tl := s[0], s[1]

	//TODO: Errors also can be thrown from these two functions
	g := Game{
		Id:      extractId(hd),
		Reveals: extractReveals(strings.Trim(tl, " ")),
	}
	return g, nil
}

func extractId(head string) int {
	id, _ := strconv.Atoi(head[5:])
	return id
}

func extractReveals(tail string) []ColorSet {
	var sets []ColorSet

	rs := strings.Split(tail, "; ")
	for _, r := range rs {
		set := ColorSet{}
		colors := strings.Split(r, ", ")
		for _, color := range colors {
			x := strings.Split(strings.Trim(color, " "), " ")
			n, _ := strconv.Atoi(x[0])

			switch x[1] {
			case "blue":
				set.Blue = n
			case "red":
				set.Red = n
			case "green":
				set.Green = n
			}
		}
		sets = append(sets, set)
	}
	return sets
}

func FewestPossibleBag(g Game) Bag {
	b := Bag{}

	for _, r := range g.Reveals {
		if r.Blue > b.Blue {
			b.Blue = r.Blue
		}
		if r.Red > b.Red {
			b.Red = r.Red
		}
		if r.Green > b.Green {
			b.Green = r.Green
		}
	}

	return Bag(b)
}
