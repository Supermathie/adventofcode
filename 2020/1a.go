package main

import (
	"supermathie.net/libadvent"
)

func day1a(inputFile string) (int, error) {
	target := 2020

	input, err := libadvent.ReadFileInts(inputFile)
	if err != nil {
		return -1, err
	}

	for _, v1 := range input {
		for _, v2 := range input {
			if v1+v2 == target {
				return v1 * v2, nil
			}
		}
	}
	return -1, adventError("target not found")
}
