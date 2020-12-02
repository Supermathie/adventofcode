package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

/*
	The policy (before the : ) describes the minimum and maximum number of times
	the character that follows must appear in the password.
*/
func day2aIsValidPassword(pass string) bool {
	// 1-7 h: hhlnhfhzxhhphhdhh
	r := regexp.MustCompile(`^(\d+)-(\d+)\s+(.):\s+(.*)$`)
	result := r.FindStringSubmatch(pass)
	if result == nil {
		log.Fatalf("regex did not match string:\n%s", pass)
	}
	min, _ := strconv.Atoi(result[1])
	max, _ := strconv.Atoi(result[2])
	letter := result[3]
	password := result[4]
	letterCount := strings.Count(password, letter)
	return min <= letterCount && letterCount <= max
}

func day2a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	var count int
	for line := range c {
		if day2aIsValidPassword(line) {
			count++
		}
	}

	return count, nil
}
