package main

import (
	"sort"

	"supermathie.net/libadvent"
)

func day5b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	seatNums := day5GetSeats(c)
	sort.Slice(seatNums, func(i, j int) bool { return seatNums[i] < seatNums[j] })
	for i, seatNum := range seatNums {
		if seatNums[i+1] == seatNum+2 {
			return seatNum + 1, nil
		}
	}

	return -1, nil
}
