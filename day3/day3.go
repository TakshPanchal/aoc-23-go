package day3

import (
	"regexp"
	"strconv"
)

func isSymbol(r byte) bool {
	return r != '.' && !(r >= '0' && r <= '9')
}

func isValidIndex(index int, s string) bool {
	return index >= 0 && index < len(s)
}

func FindPartNumbers(schema []string) []int {
	nums := []int{}
	r, _ := regexp.Compile(`\d+`)
	for i, s := range schema {
		ns := r.FindAllIndex([]byte(s), -1)
		for _, n := range ns {
			flag := isPart(i, schema, n, s)
			if flag {
				num, err := strconv.Atoi(s[n[0]:n[1]])
				if err != nil {
					panic(err)
				}
				nums = append(nums, num)
			}
		}
	}

	return nums
}

func isPart(i int, schema []string, n []int, s string) bool {
	flag := false
	// checking for symbol in upper row
	if i-1 > 0 {
		for j := n[0] - 1; j <= n[1]; j++ {
			if isValidIndex(j, schema[i-1]) && isSymbol(schema[i-1][j]) {
				flag = true
				break
			}
		}
	}
	// checking for symbol in lower row
	if i+1 < len(schema) && !flag {
		for j := n[0] - 1; j <= n[1]; j++ {
			if isValidIndex(j, schema[i+1]) && isSymbol(schema[i+1][j]) {
				flag = true
				break
			}
		}
	}

	if n[0]-1 >= 0 && isSymbol(s[n[0]-1]) {
		flag = true
	}

	if n[1] < len(s) && isSymbol(s[n[1]]) {
		flag = true
	}

	return flag
}
