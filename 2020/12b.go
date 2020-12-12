package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"supermathie.net/libadvent"
)

func day12b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	x := 0
	y := 0
	dx := 10
	dy := 1
	for dir := range c {
		action := dir[0]
		val, _ := strconv.Atoi(dir[1:])
		switch action {
		case 'N':
			dy += val
		case 'S':
			dy -= val
		case 'E':
			dx += val
		case 'W':
			dx -= val
		case 'L':
			switch val {
			case 90:
				dx, dy = -dy, dx
			case 180:
				dx, dy = -dx, -dy
			case 270:
				dx, dy = dy, -dx
			}
		case 'R':
			switch val {
			case 270:
				dx, dy = -dy, dx
			case 180:
				dx, dy = -dx, -dy
			case 90:
				dx, dy = dy, -dx
			}
		case 'F':
			x += dx * val
			y += dy * val
		default:
			log.Fatalf("bad action %v", action)
		}
	}
	fmt.Printf("%d, %d\n", x, y)

	return int(math.Abs(float64(x)) + math.Abs(float64(y))), nil
}
