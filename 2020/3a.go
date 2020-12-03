package main

import (
	"supermathie.net/libadvent"
)

func day3a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	posX := 0
	treesEncountered := 0
	<-c // we start on the first line of the input, discard it

	for line := range c {
		posX = (posX + 3) % len(line)
		if line[posX] == '#' {
			treesEncountered++
		}
	}

	return treesEncountered, nil
}
