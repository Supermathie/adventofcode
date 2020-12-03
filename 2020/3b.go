package main

import (
	"supermathie.net/libadvent"
)

func day3bCheckSlope(geography []string, xSlope, ySlope int) (treesEncountered int) {
	xPos := 0
	for yPos := 0; yPos < len(geography); yPos += ySlope {
		if geography[yPos][xPos] == '#' {
			treesEncountered++
		}
		xPos = (xPos + xSlope) % len(geography[0])
	}
	return
}

func day3b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	geography := libadvent.ChanToSlice(c)

	treesEncountered := day3bCheckSlope(geography, 1, 1)
	treesEncountered *= day3bCheckSlope(geography, 3, 1)
	treesEncountered *= day3bCheckSlope(geography, 5, 1)
	treesEncountered *= day3bCheckSlope(geography, 7, 1)
	treesEncountered *= day3bCheckSlope(geography, 1, 2)

	return treesEncountered, nil
}
