package libadvent

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRFLEmpty(t *testing.T) {
	name := "empty"
	c, err := ReadFileLines(fmt.Sprintf("testdata/%s", name))
	want := make([]string, 0)
	result := ChanToSlice(c)
	if !reflect.DeepEqual(result, want) || err != nil {
		t.Fatalf("ReadFileLines(%s) → (%v, %v), expected %v", name, result, err, want)
	}
}

func TestRFLSingle(t *testing.T) {
	name := "single"
	c, err := ReadFileLines(fmt.Sprintf("testdata/%s", name))
	want := []string{"1"}
	result := ChanToSlice(c)
	if !reflect.DeepEqual(result, want) || err != nil {
		t.Fatalf("ReadFileLines(%s) → (%v, %v), expected %v", name, result, err, want)
	}
}

func TestRFLMany(t *testing.T) {
	name := "many"
	c, err := ReadFileLines(fmt.Sprintf("testdata/%s", name))
	want := []string{"1", "2", "3", "4", "5"}
	result := ChanToSlice(c)
	if !reflect.DeepEqual(result, want) || err != nil {
		t.Fatalf("ReadFileLines(%s) → (%v, %v), expected %v", name, result, err, want)
	}
}
