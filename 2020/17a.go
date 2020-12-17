package main

import (
	"fmt"
	"log"

	"supermathie.net/libadvent"
)

type point3 struct {
	x, y, z int
}

func neighbours3(point point3) (n [27]point3) {
	// we are *deliberately* counting ourselves
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				n[(x+1)*9+(y+1)*3+(z+1)] = point3{point.x + x, point.y + y, point.z + z}
			}
		}
	}
	return
}

func day17simulate3(state map[point3]bool) (newState map[point3]bool) {
	neighbourCount := make(map[point3]int)
	newState = make(map[point3]bool)
	for point := range state {
		for _, neighbour := range neighbours3(point) {
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

func printState3(state map[point3]bool) {
	minX, maxX, minY, maxY, minZ, maxZ := 0, 0, 0, 0, 0, 0
	for point := range state {
		minX, maxX = point.x, point.x
		minY, maxY = point.y, point.y
		minZ, maxZ = point.z, point.z
		break
	}
	for point := range state {
		minX = libadvent.Min([]int{point.x, minX})
		minY = libadvent.Min([]int{point.y, minY})
		minZ = libadvent.Min([]int{point.z, minZ})
		maxX = libadvent.Max([]int{point.x, maxX})
		maxY = libadvent.Max([]int{point.y, maxY})
		maxZ = libadvent.Max([]int{point.z, maxZ})
	}
	for z := minZ; z <= maxZ; z++ {
		fmt.Printf("z=%v, maxY=%v, minX=%v\n", z, maxY, minX)
		for y := maxY; y >= minY; y-- {
			for x := minX; x <= maxX; x++ {
				if _, present := state[point3{x, y, z}]; present {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func day17a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	state := map[point3]bool{}

	{
		x, y, z := 0, 0, 0
		for line := range c {
			for x = 0; x < len(line); x++ {
				switch line[x] {
				case '.':
					// off, do nothing
				case '#':
					state[point3{x, y, z}] = true
				default:
					log.Fatalf("invalid char")
				}
			}
			y--
		}
	}

	for t := 0; t < 6; t++ {
		fmt.Printf("t=%v, count=%v\n", t, len(state))
		// printState3(state)
		state = day17simulate3(state)
	}

	return len(state), nil
}
