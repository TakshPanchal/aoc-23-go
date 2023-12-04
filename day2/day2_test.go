package day2_test

import (
	"aoc_23/day2"
	"reflect"
	"testing"
)

type testCase struct {
	l              string
	game           day2.Game
	possible       bool
	bag            day2.Bag
	minPossibleBag day2.Bag
}

var cases = []testCase{
	{
		l: "Game 1: 3 blue, 4 red",
		game: day2.Game{
			Id: 1,
			Reveals: []day2.ColorSet{
				{Blue: 3, Red: 4},
			},
		},
		bag:      day2.Bag{Blue: 3, Red: 4},
		possible: true,
	},
	{
		l: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		game: day2.Game{
			Id: 5,
			Reveals: []day2.ColorSet{
				{Red: 6, Blue: 1, Green: 3},
				{Blue: 2, Red: 1, Green: 2},
			},
		},
		bag:            day2.Bag{Blue: 3, Red: 4},
		minPossibleBag: day2.Bag{Red: 6, Green: 3, Blue: 2},
		possible:       true,
	},
}

func TestExtractGame(t *testing.T) {
	type ts struct {
		name string
		want day2.Game
		l    string
	}

	tt := []ts{
		{
			name: "Simple Game",
			l:    cases[0].l,
			want: cases[0].game,
		},
		{
			name: "two_reveals_multi_colors",
			l:    cases[1].l,
			want: cases[1].game,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := day2.ExtractGame(tc.l)

			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got: %+v but, Want: %+v", got, tc.want)
			}
		})
	}
}

func TestIsPossible(t *testing.T) {
	type ts struct {
		name string
		game day2.Game
		bag  day2.Bag
		want bool
	}

	tt := []ts{
		{
			name: "basic",
			game: cases[0].game,
			bag:  cases[0].bag,
			want: cases[0].possible,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := day2.IsPossible(tc.bag, tc.game)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got: %+v but, Want: %+v", got, tc.want)
			}
		})
	}
}

// fewest number of cubes of each color that could have been in the bag to make the game possible
func TestFewestPossibleBag(t *testing.T) {
	game := cases[1].game

	got := day2.FewestPossibleBag(game)
	want := cases[1].minPossibleBag

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %+v but, Want: %+v", got, want)
	}
}
