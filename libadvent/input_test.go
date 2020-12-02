package libadvent

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestEmptyInput(t *testing.T) {
	name := "empty"
	result, err := ReadFileInts(fmt.Sprintf("testdata/%s", name))
	want := make([]int, 0)
	if !reflect.DeepEqual(result, want) || err != nil {
		t.Fatalf("ReadFileInts(%s) → (%d, %v), expected %v", name, result, err, want)
	}
}

func TestSingleInput(t *testing.T) {
	name := "single"
	result, err := ReadFileInts(fmt.Sprintf("testdata/%s", name))
	want := []int{1}
	if !reflect.DeepEqual(result, want) || err != nil {
		t.Fatalf("ReadFileInts(%s) → (%d, %v), expected %v", name, result, err, want)
	}
}

func TestManyInput(t *testing.T) {
	name := "many"
	result, err := ReadFileInts(fmt.Sprintf("testdata/%s", name))
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, want) || err != nil {
		t.Fatalf("ReadFileInts(%s) → (%d, %v), expected %v", name, result, err, want)
	}
}

func TestManyNoEol(t *testing.T) {
	name := "manynoeol"
	result, err := ReadFileInts(fmt.Sprintf("testdata/%s", name))
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, want) || err != nil {
		t.Fatalf("ReadFileInts(%s) → (%d, %v), expected %v", name, result, err, want)
	}
}

func TestBadFile(t *testing.T) {
	name := "doesnotexist"
	result, err := ReadFileInts(fmt.Sprintf("testdata/%s", name))
	want := "os.PathError"
	if _, ok := err.(*os.PathError); !ok {
		t.Fatalf("ReadFileInts(%s) → (%d, %v), expected %v", name, result, err, want)
	}
}

func TestBadData(t *testing.T) {
	name := "baddata"
	result, err := ReadFileInts(fmt.Sprintf("testdata/%s", name))
	want := "strconv.NumError"
	if _, ok := err.(*strconv.NumError); !ok {
		t.Fatalf("ReadFileInts(%s) → (%d, %v), expected %v", name, result, err, want)
	}
}
