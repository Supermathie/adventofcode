package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"supermathie.net/libadvent"
)

func look(seats [][]rune, x, y, dir int) rune {
	if dir == 4 { // ourself
		return '.'
	}
	rX := len(seats[0])
	rY := len(seats)

	dX := dir%3 - 1
	dY := dir/3 - 1

	x += dX
	y += dY
	for x >= 0 && y >= 0 && x < rX && y < rY {
		if seats[y][x] != '.' {
			return seats[y][x]
		}
		x += dX
		y += dY
	}
	return '.' // out of bounds
}

func applyRoundB(seats [][]rune) [][]rune {
	rX := len(seats[0])
	rY := len(seats)
	newSeats := make([][]rune, rY)
	for y := 0; y < rY; y++ {
		newSeats[y] = make([]rune, rX)
		for x := 0; x < rX; x++ {
			if seats[y][x] == '.' {
				newSeats[y][x] = '.'
				continue
			}
			occupiedSeen := 0
			for i := 0; i < 9; i++ {
				if look(seats, x, y, i) == '#' {
					occupiedSeen++
				}
			}
			switch seats[y][x] {
			case 'L':
				if occupiedSeen == 0 {
					newSeats[y][x] = '#'
				} else {
					newSeats[y][x] = 'L'
				}
			case '#':
				if occupiedSeen >= 5 {
					newSeats[y][x] = 'L'
				} else {
					newSeats[y][x] = '#'
				}
			default:
				log.Fatalf("bad seat %v", seats[y][x])
			}
		}
	}
	return newSeats
}

func printSeats(seats [][]rune) {
	for y := 0; y < len(seats); y++ {
		fmt.Println(string(seats[y]))
	}
	fmt.Println()
}

func day11b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	seats := make([][]rune, 0)
	for i := range c {
		seats = append(seats, []rune(i))
	}

	printSeats(seats)
	for newSeats := applyRoundB(seats); !reflect.DeepEqual(seats, newSeats); {
		seats = newSeats
		printSeats(seats)
		newSeats = applyRoundB(seats)
		time.Sleep(time.Millisecond * 100)
	}

	printSeats(seats)

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
