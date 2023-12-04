package day1

import (
	"strconv"
	"strings"
)

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

var mp = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

// returns calibration digit
func Calibration(l string) int {
	fi, li := -1, -1
	var fv, lv rune

	for i, c := range l {
		if isDigit(c) {
			if fi == -1 {
				fi, li = i, i
				fv, lv = c, c
			} else {
				li = i
				lv = c
			}
		}
	}

	for k, v := range mp {
		is := strings.Index(strings.ToLower(l), k)
		ie := strings.LastIndex(strings.ToLower(l), k)

		if is != -1 {
			if fi == -1 || is < fi {
				fv = v
				fi = is
			}

			if ie > li {
				li = ie
				lv = v
			}
		}
	}

	cv, err := strconv.Atoi(string(fv) + string(lv))
	if err != nil {
		panic(err)
	}

	return cv
}
