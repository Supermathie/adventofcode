package main

import (
	"fmt"
	"testing"
)

func Test15aCase1(t *testing.T) {
	data := "15_1"
	result, err := day15a(fmt.Sprintf("testdata/%s", data))
	want := 165
	if result != want || err != nil {
		t.Fatalf("day15a(%s) → (%d, %v), expected %v", data, result, err, want)
	}
}

func Test15aCase2(t *testing.T) {
	data := "15_2"
	result, err := day15a(fmt.Sprintf("testdata/%s", data))
	want := 137438953471

	if result != want || err != nil {
		t.Fatalf("day15a(%s) → (%d, %v), expected %v", data, result, err, want)
	}
}
