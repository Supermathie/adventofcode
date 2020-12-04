package main

import (
	"testing"
)

func Test2aCase1(t *testing.T) {
	// 1-3 a: abcde is valid
	data := "1-3 a: abcde"
	result := day2aIsValidPassword(data)
	want := true
	if result != want {
		t.Fatalf("day2a(%s) → (%v), expected %v", data, result, want)
	}
}

func Test2aCase2(t *testing.T) {
	// 1-3 b: cdefg is invalid: not enough b
	data := "1-3 b: cdefg"
	result := day2aIsValidPassword(data)
	want := false
	if result != want {
		t.Fatalf("day2a(%s) → (%v), expected %v", data, result, want)
	}
}

func Test2aCase3(t *testing.T) {
	// 2-9 c: ccccccccc is valid
	data := "2-9 c: ccccccccc"
	result := day2aIsValidPassword(data)
	want := true
	if result != want {
		t.Fatalf("day2a(%s) → (%v), expected %v", data, result, want)
	}
}
