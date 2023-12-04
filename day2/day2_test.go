package day2_test

import (
	"aoc_23/day2"
	"reflect"
	"testing"
)

func TestIsPossible(t *testing.T) {
	type ts struct {
		name string
		game day2.Game
		bag  day2.Bag
		want bool
	}

	tt := []ts{
		{name: "basic", game: day2.Game{Id: 1, Reveals: []day2.ColorSet{
			{Blue: 3, Red: 4},
		}}, bag: day2.Bag{
			Blue: 3, Red: 4,
		}, want: true},
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

func TestExtractGame(t *testing.T) {
	type ts struct {
		name string
		want day2.Game
		l    string
	}

	tt := []ts{
		{
			name: "Simple Game",
			l:    "Game 1: 3 blue, 4 red",
			want: day2.Game{Id: 1, Reveals: []day2.ColorSet{
				{Blue: 3, Red: 4},
			}},
		},
		{
			name: "two_reveals_multi_colors",
			l:    "Game 5: 60 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: day2.Game{
				Id: 5, Reveals: []day2.ColorSet{
					{Red: 60, Blue: 1, Green: 3},
					{Blue: 2, Red: 1, Green: 2},
				},
			},
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
