package main

import (
	"log"

	"supermathie.net/libadvent"
)

func day9b(inputFile string) (int, error) {
	allData, err := libadvent.ReadFileInts(inputFile)
	if err != nil {
		return -1, err
	}

	target, _ := day9a(inputFile)
	i, err := libadvent.IndexOf(allData, target)
	if err != nil {
		log.Fatalf("Whaaaaaaat? %v is not in the input? %s", target, err)
	}
	for start := i - 2; start > 0; start-- {
		for end := i - 1; end > start; end-- {
			sum := libadvent.Sum(allData[start : end+1])
			if sum == target {
				return (libadvent.Min(allData[start:end+1]) + libadvent.Max(allData[start:end+1])), nil
			}
			if sum < target {
				break
			}
		}
	}
	return -1, nil
}
