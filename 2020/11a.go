package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"supermathie.net/libadvent"
)

func applyRound(seats []string) []string {
	rX := len(seats[0])
	rY := len(seats)
	newSeats := make([]string, 0)
	for y := 0; y < rY; y++ {
		newSeatLine := make([]string, 0)
		for x := 0; x < rX; x++ {
			if seats[y][x] == '.' {
				newSeatLine = append(newSeatLine, ".")
				continue
			}
			occupiedAdjacent := 0
			if y > 0 {
				occupiedAdjacent += strings.Count(seats[y-1][libadvent.Max([]int{0, x - 1}):libadvent.Min([]int{rX, x + 2})], "#")
			}
			occupiedAdjacent += strings.Count(seats[y][libadvent.Max([]int{0, x - 1}):libadvent.Min([]int{rX, x + 2})], "#")
			if y < rY-1 {
				occupiedAdjacent += strings.Count(seats[y+1][libadvent.Max([]int{0, x - 1}):libadvent.Min([]int{rX, x + 2})], "#")
			}
			switch seats[y][x] {
			case 'L':
				if occupiedAdjacent == 0 {
					newSeatLine = append(newSeatLine, "#")
				} else {
					newSeatLine = append(newSeatLine, "L")
				}
			case '#':
				if occupiedAdjacent-1 >= 4 { // we counted ourself
					newSeatLine = append(newSeatLine, "L")
					//newSeatLine = append(newSeatLine, fmt.Sprintf("%d", occupiedAdjacent-1))
				} else {
					newSeatLine = append(newSeatLine, "#")
					//newSeatLine = append(newSeatLine, fmt.Sprintf("%d", occupiedAdjacent-1))
				}
			default:
				log.Fatalf("bad seat %v", seats[y][x])
			}
		}
		newSeats = append(newSeats, strings.Join(newSeatLine, ""))
	}
	return newSeats
}

func day11a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	seats := make([]string, 0)
	for i := range c {
		seats = append(seats, i)
	}

	fmt.Printf("%v\n\n", strings.Join(seats, "\n"))
	for newSeats := applyRound(seats); !reflect.DeepEqual(seats, newSeats); {
		seats = newSeats
		fmt.Printf("%v\n\n", strings.Join(seats, "\n"))
		newSeats = applyRound(seats)
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%v\n", strings.Join(seats, "\n"))
	return strings.Count(strings.Join(seats, ""), "#"), nil
}
