package main

import (
	"log"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

func day13b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	<-c // discard the first line
	buses := map[uint64]uint64{}
	for i, bus := range strings.Split(<-c, ",") {
		if bus == "x" {
			continue
		}
		busNum, _ := strconv.Atoi(bus)
		buses[uint64(i)] = uint64(busNum)
	}

	time, found := buses[0]
	if !found {
		log.Fatalf("I expected a bus at slot 0")
	}
	delete(buses, 0)
	timeStep := uint64(time)

	for departureDelta, busNum := range buses {
		for (time+departureDelta)%busNum != 0 {
			time += timeStep
		}
		timeStep = libadvent.LCM(busNum, timeStep)
	}

	return int(time), nil
}
