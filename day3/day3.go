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

type numIndex struct {
	start, end int
}

func findAllNumIndexes(schema []string) (idx [][]numIndex) {
	r, _ := regexp.Compile(`\d+`)
	for _, s := range schema {
		row := []numIndex{}
		ns := r.FindAllIndex([]byte(s), -1)
		for _, n := range ns {
			row = append(row, numIndex{start: n[0], end: n[1] - 1})
		}
		idx = append(idx, row)
	}
	return
}

func FindPartNumbers(schema []string) []int {
	numIdx := findAllNumIndexes(schema)
	nums := []int{}

	for r, row := range numIdx {
		for _, ni := range row {
			flag := isPart(r, schema, []int{ni.start, ni.end + 1}, schema[r])
			if flag {
				num, err := strconv.Atoi(schema[r][ni.start : ni.end+1])
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

type Gear struct {
	N1, N2 int
}

func FindGears(schema []string) (gears []Gear) {
	numIdx := findAllNumIndexes(schema)
	r, _ := regexp.Compile(`\*`)
	for rId, s := range schema {
		starts := r.FindAllIndex([]byte(s), -1)
		for _, star := range starts {
			adjPart := 0
			n1, n2 := 0, 0

			// check adj in the same line
			for _, num := range numIdx[rId] {
				if star[0] == num.end+1 || star[0] == num.start-1 {
					adjPart++
					if adjPart == 1 {
						n1, _ = strconv.Atoi(s[num.start : num.end+1])
					} else if adjPart == 2 {
						n2, _ = strconv.Atoi(s[num.start : num.end+1])
						gears = append(gears, Gear{n1, n2})
					}
				}
			}

			// check adj in upper row
			if rId-1 >= 0 {
				for _, num := range numIdx[rId-1] {
					if indexOverlap(num, star[0]) {
						adjPart++
						if adjPart == 1 {
							n1, _ = strconv.Atoi(schema[rId-1][num.start : num.end+1])
						} else if adjPart == 2 {
							n2, _ = strconv.Atoi(schema[rId-1][num.start : num.end+1])
							gears = append(gears, Gear{n1, n2})
						}
					}
				}
			}

			if rId+1 < len(schema) {
				for _, num := range numIdx[rId+1] {
					if indexOverlap(num, star[0]) {
						adjPart++
						if adjPart == 1 {
							n1, _ = strconv.Atoi(schema[rId+1][num.start : num.end+1])
						} else if adjPart == 2 {
							n2, _ = strconv.Atoi(schema[rId+1][num.start : num.end+1])
							gears = append(gears, Gear{n1, n2})
						}
					}
				}
			}
		}
	}
	return gears
}

func indexOverlap(p numIndex, i int) bool {
	return ((p.start <= i && i <= p.end) || p.start == i || p.end == i) ||
		((p.start <= i-1 && i-1 <= p.end) || p.start == i-1 || p.end == i-1) ||
		((p.start <= i+1 && i+1 <= p.end) || p.start == i+1 || p.end == i+1)
}
