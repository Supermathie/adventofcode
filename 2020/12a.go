package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"supermathie.net/libadvent"
)

func day12a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	x := 0
	y := 0
	heading := 90
	for dir := range c {
		action := dir[0]
		val, _ := strconv.Atoi(dir[1:])
		switch action {
		case 'N':
			y += val
		case 'S':
			y -= val
		case 'E':
			x += val
		case 'W':
			x -= val
		case 'L':
			heading = (heading - val + 360) % 360
		case 'R':
			heading = (heading + val + 360) % 360
		case 'F':
			switch heading {
			case 0:
				y += val
			case 180:
				y -= val
			case 90:
				x += val
			case 270:
				x -= val
			default:
				log.Fatalf("bad heading %d", heading)
			}
		default:
			log.Fatalf("bad action %v", action)
		}
	}
	fmt.Printf("%d, %d\n", x, y)

	return 0, errors.New("not implemented")
}
