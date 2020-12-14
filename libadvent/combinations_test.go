package libadvent

import (
	"reflect"
	"testing"
)

func TestAllCombinations0(t *testing.T) {
	options := []int{}
	result := ChanToSliceSI(AllCombinations(options))
	want := [][]int{
		[]int{},
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("TestAllCombinations0(%v) → (%v), expected %v", options, result, want)
	}
}

func TestAllCombinations1(t *testing.T) {
	options := []int{1}
	result := ChanToSliceSI(AllCombinations(options))
	want := [][]int{
		[]int{},
		[]int{1},
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("TestAllCombinations1(%v) → (%v), expected %v", options, result, want)
	}
}

func TestAllCombinations3(t *testing.T) {
	options := []int{1, 2, 3}
	result := ChanToSliceSI(AllCombinations(options))
	want := [][]int{
		[]int{},
		[]int{1},
		[]int{2},
		[]int{3},
		[]int{1, 2},
		[]int{1, 3},
		[]int{2, 3},
		[]int{1, 2, 3},
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("TestAllCombinations3(%v) → (%v), expected %v", options, result, want)
	}
}

func TestCombinations0(t *testing.T) {
	options := []int{1, 2, 3}
	result := ChanToSliceSI(Combinations(options, 0))
	want := [][]int{
		[]int{},
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("TestCombinations0(%v) → (%v), expected %v", options, result, want)
	}
}

func TestCombinations1(t *testing.T) {
	options := []int{1, 2, 3}
	result := ChanToSliceSI(Combinations(options, 1))
	want := [][]int{
		{1},
		{2},
		{3},
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("TestCombinations1(%v) → (%v), expected %v", options, result, want)
	}
}

func TestCombinations2(t *testing.T) {
	options := []int{1, 2, 3}
	result := ChanToSliceSI(Combinations(options, 2))
	want := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("TestCombinations2(%v) → (%v), expected %v", options, result, want)
	}
}
