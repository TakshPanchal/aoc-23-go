package utils

import (
	"io"
	"io/fs"
	"strings"
)

func ReadLines(fSys fs.FS, fname string) ([]string, error) {
	f, err := fSys.Open(fname)
	if err != nil {
		return []string{}, nil
	}
	defer f.Close()

	fData, err := io.ReadAll(f)
	if err != nil {
		return []string{}, nil
	}

	return strings.Split(string(fData), "\n"), nil
}
