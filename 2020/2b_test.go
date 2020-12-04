package main

import (
	"testing"
)

func Test2bCase1(t *testing.T) {
	// 1-3 a: abcde is valid: position 1 contains a and position 3 does not.
	data := "1-3 a: abcde"
	result := day2bIsValidPassword(data)
	want := true
	if result != want {
		t.Fatalf("day2b(%s) → (%v), expected %v", data, result, want)
	}
}

func Test2bCase2(t *testing.T) {
	// 1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
	data := "1-3 b: cdefg"
	result := day2bIsValidPassword(data)
	want := false
	if result != want {
		t.Fatalf("day2b(%s) → (%v), expected %v", data, result, want)
	}
}

func Test2bCase3(t *testing.T) {
	// 2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
	data := "2-9 c: ccccccccc"
	result := day2bIsValidPassword(data)
	want := false
	if result != want {
		t.Fatalf("day2b(%s) → (%v), expected %v", data, result, want)
	}
}
