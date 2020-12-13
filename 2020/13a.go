package main

import (
	"errors"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

func day13a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	startTime, _ := strconv.Atoi(<-c)

	buses := []int{}
	for _, bus := range strings.Split(<-c, ",") {
		if bus == "x" {
			continue
		}
		busNum, _ := strconv.Atoi(bus)
		buses = append(buses, busNum)
	}
	for timeDelta := 0; true; timeDelta++ {
		for _, bus := range buses {
			if (startTime+timeDelta)%bus == 0 {
				return timeDelta * bus, nil
			}
		}
	}

	return 0, errors.New("this cannot possibly happen")
}
