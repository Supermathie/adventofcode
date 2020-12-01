package main

import (
	"log"

	"supermathie.net/libadvent"
)

func day1b(inputFile string) int {
	target := 2020

	input, err := libadvent.ReadFileInts(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, v1 := range input {
		for _, v2 := range input {
			for _, v3 := range input {
				if v1+v2+v3 == target {
					return v1 * v2 * v3
				}
			}
		}
	}
	log.Fatal("target not found")
	return -1
}
