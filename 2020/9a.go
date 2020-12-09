package main

import (
	"supermathie.net/libadvent"
)

func day9a(inputFile string) (int, error) {
	allData, err := libadvent.ReadFileInts(inputFile)
	if err != nil {
		return -1, err
	}

	for i := 25; i < len(allData); i++ {
		if libadvent.FindCombinationTotal(allData[i-25:i], 2, allData[i]) == nil {
			return allData[i], nil
		}
	}
	return -1, nil
}
