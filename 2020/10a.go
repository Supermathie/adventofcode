package main

import (
	"fmt"
	"sort"

	"supermathie.net/libadvent"
)

func day10a(inputFile string) (int, error) {
	input, err := libadvent.ReadFileInts(inputFile)
	if err != nil {
		return -1, err
	}
	data := make([]int, len(input)+1)
	copy(data[1:], input) // leave a 0 at the beginning to represent the seat
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	data = append(data, data[len(data)-1]+3) // our device is 3 higher than the highest adaptor

	num1 := 0
	num3 := 0

	fmt.Println(data)

	for i := 0; i < len(data)-1; i++ {
		if data[i]+1 == data[i+1] {
			num1++
		}
		if data[i]+3 == data[i+1] {
			num3++
		}
	}
	return num1 * num3, nil
}
