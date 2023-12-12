package day11

import (
	"math"
	"regexp"
	"strings"
)

type Image []string

type Galaxy struct {
	X, Y int
}

func ExpandImage(img Image) Image {
	cols := getColumns(img)
	rIdx := getNoGalaxyIndexAll(img)
	cInx := getNoGalaxyIndexAll(cols)

	expImg := Image{}
	rPtr := 0
	for r := range img {
		cPtr := 0
		var sb strings.Builder
		for c := range img[r] {
			sb.WriteRune(rune(img[r][c]))
			if cPtr < len(cInx) && c == cInx[cPtr] {
				sb.WriteRune('.')
				cPtr++
			}
		}
		expImg = append(expImg, sb.String())

		if rPtr < len(rIdx) && r == rIdx[rPtr] {
			newRow := strings.Repeat(".", sb.Len())
			expImg = append(expImg, newRow)
			rPtr++
		}
	}
	return expImg
}

func getColumns(img Image) (cols []string) {

	for c := range img[0] {
		col := ""
		for r := range img {
			col += string(img[r][c])
		}

		cols = append(cols, col)
	}
	return
}

func getNoGalaxyIndexAll(lists []string) (idx []int) {
	r, _ := regexp.Compile("#")

	for i, l := range lists {
		if !r.MatchString(l) {
			idx = append(idx, i)
		}
	}
	return
}

func FindClosestDistance(g1, g2 Galaxy) int {
	return int(math.Abs(float64(g1.X)-float64(g2.X)) +
		math.Abs(float64(g1.Y)-float64(g2.Y)))
}

func FindGalaxies(img Image) (gs []Galaxy) {
	r, _ := regexp.Compile("#")

	for rI := range img {
		idx := r.FindAllIndex([]byte(img[rI]), -1)

		for _, i := range idx {
			gs = append(gs, Galaxy{X: rI + 1, Y: i[1]})

		}
	}
	return
}

func GetEmptyRowColIdx(img Image) (idx [2][]int) {
	cols := getColumns(img)
	rIdx := getNoGalaxyIndexAll(img)
	cInx := getNoGalaxyIndexAll(cols)
	idx[0] = rIdx
	idx[1] = cInx
	return
}
