package main

import (
	"fmt"
	"testing"
)

func Test1bCase1(t *testing.T) {
	data := "1_1"
	result, err := day1b(fmt.Sprintf("testdata/%s", data))
	want := 241861950
	if result != want || err != nil {
		t.Fatalf("day1a(%s) → (%d, %v), expected %v", data, result, err, want)
	}
}

func Test1bCase2(t *testing.T) {
	data := "1_2"
	result, err := day1b(fmt.Sprintf("testdata/%s", data))

	if err == nil {
		t.Fatalf("day1a(%s) → (%d, %v), expected error", data, result, err)
	}
}
