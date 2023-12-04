package utils_test

import (
	"aoc_23/utils"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestRead(t *testing.T) {
	// os.DirFS("..")
	fs := fstest.MapFS{
		"sample.txt": {Data: []byte("line1\nline2")},
	}

	got, err := utils.ReadLines(fs, "sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"line1",
		"line2",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v but want: %v", got, want)
	}
}
