package utils

import (
	"os"
	"strings"
)

func ReadLines(fname string) []string {
	fData, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(fData), "\n")
}
