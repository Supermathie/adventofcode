package libadvent

import (
	"testing"
)

func TestLCM21(t *testing.T) {
	var a, b uint64
	a, b = 3, 7
	result := LCM(a, b)
	want := uint64(21)
	if result != want {
		t.Fatalf("LCM(%d, %d) → (%d), expected %d", a, b, result, want)
	}
}

func TestLCM24(t *testing.T) {
	var a, b uint64
	a, b = 8, 6
	result := LCM(a, b)
	want := uint64(24)
	if result != want {
		t.Fatalf("LCM(%d, %d) → (%d), expected %d", a, b, result, want)
	}
}

func TestGCD1(t *testing.T) {
	var a, b uint64
	a, b = 3, 7
	result := GCD(a, b)
	want := uint64(1)
	if result != want {
		t.Fatalf("GCD(%d, %d) → (%d), expected %d", a, b, result, want)
	}
}

func TestGCD5(t *testing.T) {
	var a, b uint64
	a, b = 10, 25
	result := GCD(a, b)
	want := uint64(5)
	if result != want {
		t.Fatalf("GCD(%d, %d) → (%d), expected %d", a, b, result, want)
	}
}
