package main

import (
	"fmt"
	"testing"
)

func Test14aCase1(t *testing.T) {
	data := "14_1"
	result, err := day14a(fmt.Sprintf("testdata/%s", data))
	want := 165
	if result != want || err != nil {
		t.Fatalf("day14a(%s) → (%d, %v), expected %v", data, result, err, want)
	}
}

func Test14aCase2(t *testing.T) {
	data := "14_2"
	result, err := day14a(fmt.Sprintf("testdata/%s", data))
	want := 137438953471

	if result != want || err != nil {
		t.Fatalf("day14a(%s) → (%d, %v), expected %v", data, result, err, want)
	}
}
