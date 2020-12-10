package main

import (
	"fmt"
	"sort"

	"supermathie.net/libadvent"
)

func validPath(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i]+3 < a[i+1] {
			// fmt.Printf("-:%v\n", a)
			return false
		}
	}
	// fmt.Printf("+:%v\n", a)
	return true
}

func prepend(a []int, b int) []int {
	new := make([]int, len(a)+1)
	new[0] = b
	copy(new[1:], a)
	return new
}

func validCombinations(a []int) int {
	total := 1
	for i := 1; i < len(a)-2; i++ {
		for combination := range libadvent.Combinations(a[1:(len(a)-1)], i) {
			if validPath(prepend(append(combination, a[len(a)-1]), a[0])) {
				total++
			}
		}
	}
	return total
}

func day10b(inputFile string) (int, error) {
	input, err := libadvent.ReadFileInts(inputFile)
	if err != nil {
		return -1, err
	}
	data := make([]int, len(input)+1)
	copy(data[1:], input) // leave a 0 at the beginning to represent the seat
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	data = append(data, data[len(data)-1]+3) // our device is 3 higher than the highest adaptor
	fmt.Println(data)

	total := 1
	splitPoints := make([]int, 1)
	for i := 0; i < len(data)-1; i++ {
		if data[i]+3 == data[i+1] {
			splitPoints = append(splitPoints, i)
		}
	}

	for i := 0; i < len(splitPoints)-1; i++ {
		subA := data[splitPoints[i] : splitPoints[i+1]+1]
		subValid := validCombinations(subA)
		fmt.Printf("%v has %d\n", subA, subValid)
		total *= subValid
	}

	return total, nil
}
