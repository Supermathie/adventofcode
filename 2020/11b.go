package main

import (
	"math"

	"supermathie.net/libadvent"
)

func day11b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	seats := make([][]rune, 0)
	for i := range c {
		seats = append(seats, []rune(i))
	}

	for changed := true; changed == true; seats, changed = applyRound(seats, 5, math.MaxUint16) {
		// printSeats(seats)
	}

	totalOccupied := 0
	for y := 0; y < len(seats); y++ {
		for x := 0; x < len(seats[y]); x++ {
			if seats[y][x] == '#' {
				totalOccupied++
			}
		}
	}
	return totalOccupied, nil
}
