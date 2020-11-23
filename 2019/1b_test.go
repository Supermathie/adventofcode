package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestCalculateFuel12(t *testing.T) {
	weight := 12
	want := 2
	msg := calculateFuel(weight)
	if want != msg {
		t.Fatalf(`calculateFuel(%d) = %d, not %d`, weight, msg, want)
	}
}
