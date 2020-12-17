package main

import (
	"fmt"
	"log"

	"supermathie.net/libadvent"
)

type point4 struct {
	w, x, y, z int
}

func neighbours4(point point4) (n [81]point4) {
	// we are *deliberately* counting ourselves
	for w := -1; w <= 1; w++ {
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for z := -1; z <= 1; z++ {
					n[(w+1)*27+(x+1)*9+(y+1)*3+(z+1)] = point4{point.w + w, point.x + x, point.y + y, point.z + z}
				}
			}
		}
	}
	return
}

func day17simulate4(state map[point4]bool) (newState map[point4]bool) {
	neighbourCount := make(map[point4]int)
	newState = make(map[point4]bool)
	for point := range state {
		for _, neighbour := range neighbours4(point) {
			if _, present := neighbourCount[neighbour]; present {
				neighbourCount[neighbour]++
			} else {
				neighbourCount[neighbour] = 1
			}
		}
	}
	for point, count := range neighbourCount {
		if count == 3 || (state[point] == true && count == 4) {
			newState[point] = true
		}
	}
	return
}

func printState4(state map[point4]bool) {
	// this function left as an exercise for the reader
}

func day17b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	state := map[point4]bool{}

	{
		w, x, y, z := 0, 0, 0, 0
		for line := range c {
			for x = 0; x < len(line); x++ {
				switch line[x] {
				case '.':
					// off, do nothing
				case '#':
					state[point4{w, x, y, z}] = true
				default:
					log.Fatalf("invalid char")
				}
			}
			y--
		}
	}

	for t := 0; t < 6; t++ {
		fmt.Printf("t=%v, count=%v\n", t, len(state))
		printState4(state)
		state = day17simulate4(state)
	}

	return len(state), nil
}
