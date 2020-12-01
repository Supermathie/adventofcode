package main

import (
	"fmt"
	"testing"
)

func Test1aCase1(t *testing.T) {
	name := "1_1"
	result, err := day1a(fmt.Sprintf("testdata/%s", name))
	want := 514579
	if result != want || err != nil {
		t.Fatalf("day1a(%s) → (%d, %v), expected %v", name, result, err, want)
	}
}

func Test1aCase2(t *testing.T) {
	name := "1_2"
	result, err := day1a(fmt.Sprintf("testdata/%s", name))

	if err == nil {
		t.Fatalf("day1a(%s) → (%d, %v), expected error", name, result, err)
	}
}
