package main

import (
	"testing"
)

func Test2bCase1(t *testing.T) {
	// 1-3 a: abcde is valid: position 1 contains a and position 3 does not.
	pwent := "1-3 a: abcde"
	result := day2bIsValidPassword(pwent)
	want := true
	if result != want {
		t.Fatalf("day2b(%s) → (%v), expected %v", pwent, result, want)
	}
}

func Test2bCase2(t *testing.T) {
	// 1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
	pwent := "1-3 b: cdefg"
	result := day2bIsValidPassword(pwent)
	want := false
	if result != want {
		t.Fatalf("day2b(%s) → (%v), expected %v", pwent, result, want)
	}
}

func Test2bCase3(t *testing.T) {
	// 2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
	pwent := "2-9 c: ccccccccc"
	result := day2bIsValidPassword(pwent)
	want := false
	if result != want {
		t.Fatalf("day2b(%s) → (%v), expected %v", pwent, result, want)
	}
}
