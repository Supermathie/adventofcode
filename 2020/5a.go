package main

import (
	"log"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

func day5GetSeats(c chan string) []int {
	xlate := strings.NewReplacer(
		"B", "1",
		"F", "0",
		"R", "1",
		"L", "0",
	)
	seats := make([]int, 0)
	for seat := range c {
		seatNum, err := strconv.ParseInt(xlate.Replace(seat), 2, 0)
		if err != nil {
			log.Fatalf("error parsing %s", seat)
		}
		seats = append(seats, int(seatNum))
	}
	return seats
}

func day5a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	max := -1
	for _, seatNum := range day5GetSeats(c) {
		if seatNum > max {
			max = seatNum
		}
	}

	return max, nil
}
