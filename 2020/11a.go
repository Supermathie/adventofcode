package main

import (
	"fmt"
	"log"
	"time"

	"supermathie.net/libadvent"
)

func applyRound(seats [][]rune, abandonThreshold uint, visionDistance uint) (newSeats [][]rune, changed bool) {
	rX := len(seats[0])
	rY := len(seats)
	newSeats = make([][]rune, rY)
	changed = false
	for y := 0; y < rY; y++ {
		newSeats[y] = make([]rune, rX)
		for x := 0; x < rX; x++ {
			if seats[y][x] == '.' {
				newSeats[y][x] = '.'
				continue
			}
			occupiedSeen := uint(0)
			for i := 0; i < 9; i++ {
				if look(seats, x, y, i, visionDistance) == '#' {
					occupiedSeen++
				}
			}
			switch seats[y][x] {
			case 'L':
				if occupiedSeen == 0 {
					newSeats[y][x] = '#'
					changed = true
				} else {
					newSeats[y][x] = 'L'
				}
			case '#':
				if occupiedSeen >= abandonThreshold {
					newSeats[y][x] = 'L'
					changed = true
				} else {
					newSeats[y][x] = '#'
				}
			default:
				log.Fatalf("bad seat %v", seats[y][x])
			}
		}
	}
	return
}

func look(seats [][]rune, x, y, dir int, visionDistance uint) rune {
	if dir == 4 { // ourself
		return '.'
	}
	rX := len(seats[0])
	rY := len(seats)

	dX := dir%3 - 1
	dY := dir/3 - 1

	x += dX
	y += dY
	for dist := uint(1); x >= 0 && y >= 0 && x < rX && y < rY && dist <= visionDistance; dist++ {
		if seats[y][x] != '.' {
			return seats[y][x]
		}
		x += dX
		y += dY
	}
	return '.' // out of bounds
}

func printSeats(seats [][]rune) {
	for y := 0; y < len(seats); y++ {
		fmt.Println(string(seats[y]))
	}
	fmt.Println()
	time.Sleep(time.Millisecond * 75)
}

func day11a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	seats := make([][]rune, 0)
	for i := range c {
		seats = append(seats, []rune(i))
	}

	for changed := true; changed == true; seats, changed = applyRound(seats, 4, 1) {
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
